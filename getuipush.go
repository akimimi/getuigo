package getuigo

import (
	"errors"
	"fmt"
	"github.com/geek-go/getui"
	"github.com/satori/go.uuid"
	"time"
)

type GetuiPush struct {
	Config *GetuiConfig
}

func (g *GetuiPush) SendTransmissionByCid(cid string, payload Payload) error {

	// get auth token
	token, _ := getui.GetGeTuiToken(g.Config.AppId, g.Config.AppKey, g.Config.MasterSecret)

	// build message body
	message := getui.GetMessage()
	message.AppKey = g.Config.AppKey
	message.MsgType = getui.MsgType.Transmission

	// build transmission template by payload
	template, pushInfo, err := IGtTransmissionTemplate(payload)
	if err != nil {
		return err
	}

	pushSingleParam := &getui.PushSingleParam{
		Message:      message,
		Transmission: template,
		Cid:          cid,
		PushInfo:     pushInfo,
		RequestId:    g.RequestId(),
	}

	res, err := getui.PushSingle(g.Config.AppId, token, pushSingleParam)
	if err != nil {
		return err
	}

	defer logSinglePush(res)
	return nil
}

func (g *GetuiPush) SendTransmissionByCidList(cids []string, payload Payload) error {

	token, _ := getui.GetGeTuiToken(g.Config.AppId, g.Config.AppKey, g.Config.MasterSecret)
	message := getui.GetMessage()
	message.AppKey = g.Config.AppKey
	message.MsgType = getui.MsgType.Transmission

	template, pushInfo, err := IGtTransmissionTemplate(payload)
	if err != nil {
		return err
	}

	saveListBodyParam := &getui.SaveListBodyParam{
		Message:      message,
		Transmission: template,
		PushInfo:     pushInfo, //必须
	}

	res, err := getui.SaveListBody(g.Config.AppId, token, saveListBodyParam)
	if err != nil {
		return err
	}
	if res.Result != "ok" {
		return errors.New(fmt.Sprintf("获取contentId失败:%s,%s", res.Result, res.Desc))
	}

	taskid := res.TaskId
	pushListParam := &getui.PushListParam{
		Cid:        cids,
		Taskid:     taskid,
		NeedDetail: true,
	}
	if res2, err := getui.PushList(g.Config.AppId, token, pushListParam); err == nil {
		defer logListPush(res2)
		return nil
	} else {
		return err
	}
}

func (g *GetuiPush) SendTransmissionToAll(payload Payload, filter ...getui.AppCondition) error {
	token, _ := getui.GetGeTuiToken(g.Config.AppId, g.Config.AppKey, g.Config.MasterSecret)
	message := getui.GetMessage()
	message.AppKey = g.Config.AppKey
	message.MsgType = getui.MsgType.Transmission

	template, pushInfo, err := IGtTransmissionTemplate(payload)
	if err != nil {
		return err
	}

	conditions := g.PushAppConditions(filter...)
	pushAppParam := &getui.PushAppParam{
		Message:      message,
		Transmission: template,
		PushInfo:     pushInfo,
		Condition:    &conditions,
		RequestId:    g.RequestId(),
	}

	if res, err := getui.PushApp(g.Config.AppId, token, pushAppParam); err == nil {
		defer logAppPush(res)
		return nil
	} else {
		return err
	}
}

func (g *GetuiPush) PushAppConditions(filters ...getui.AppCondition) getui.Condition {
	isPhoneTypeSet := false
	conditions := g.MergeAppConditions(filters...)
	for _, cond := range conditions {
		if cond.Key == getui.PHONE_TYPE {
			isPhoneTypeSet = true
		}
	}
	if !isPhoneTypeSet {
		conditions = append(conditions, getui.AppCondition{
			Key:    getui.PHONE_TYPE,
			Values: []string{"ANDROID", "IOS"},
		})
	}
	return conditions
}

func (g *GetuiPush) MergeAppConditions(filters ...getui.AppCondition) getui.Condition {
	conditions := getui.Condition{}
	for _, cond := range filters {
		conditions = append(conditions, cond)
	}
	return conditions
}

func (g *GetuiPush) RequestId() (s string) {
	u2, err := uuid.NewV4()
	if err != nil {
		s = time.Now().Format("20160102150405")
		defer logUnexpected(fmt.Sprintf("uuid can not be generated: %s instead", s))
	} else {
		s = u2.String()
	}
	return
}
