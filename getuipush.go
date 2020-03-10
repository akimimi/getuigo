package getuigo

import (
	"errors"
	"fmt"
	"github.com/geek-go/getui"
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

func (g *GetuiPush) RequestId() string {
	return time.Now().Format("20160102150405")
}
