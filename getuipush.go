package getuigo

import (
	"errors"
	"fmt"
	"github.com/geek-go/getui"
	"github.com/satori/go.uuid"
	"math/rand"
	"time"
)

type GetuiPush struct {
	Config          *GetuiConfig
	Token           string
	TokenExpires    time.Time
	RequestTokenCnt int
}

func (g *GetuiPush) IsAuthTokenValid() bool {
	return g.Token != "" && g.TokenExpires.After(time.Now())
}

func (g *GetuiPush) GetAuthToken() (string, error) {
	if g.IsAuthTokenValid() { // using cache
		return g.Token, nil
	} else {
		g.RequestTokenCnt++
		token, err := getui.GetGeTuiToken(g.Config.AppId, g.Config.AppKey, g.Config.MasterSecret)
		if err != nil {
			g.Token = ""
			return "", err
		} else {
			if token == "" {
				g.Token = ""
				return "", errors.New("Token empty, auth_overlimit might happens")
			}
			g.Token = token
			g.TokenExpires = time.Now().Add(time.Hour * 12)
			return token, nil
		}
	}
}

func (g *GetuiPush) SendTransmissionByCid(cid string, payload Payload) error {
	// get auth token
	token, errToken := g.GetAuthToken()
	if errToken != nil {
		return errors.New("[GetAuthToken]" + errToken.Error())
	}

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
		RequestId:    g.RequestId(true),
	}

	res, err := getui.PushSingle(g.Config.AppId, token, pushSingleParam)
	if err != nil {
		return err
	}

	defer logSinglePush(res)
	return nil
}

func (g *GetuiPush) SendTransmissionByCidList(cids []string, payload Payload) error {
	token, errToken := g.GetAuthToken()
	if errToken != nil {
		return errors.New("[GetAuthToken]" + errToken.Error())
	}

	message := getui.GetMessage()
	message.AppKey = g.Config.AppKey
	message.MsgType = getui.MsgType.Transmission

	template, pushInfo, err := IGtTransmissionTemplate(payload)
	if err != nil {
		return errors.New("[TransTemplate]" + err.Error())
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
		return errors.New(fmt.Sprintf("[ContentId]%s,%s,%s", res.Result, res.Desc, token))
	}

	taskId := res.TaskId
	pushListParam := &getui.PushListParam{
		Cid:        cids,
		Taskid:     taskId,
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
	token, errToken := g.GetAuthToken()
	if errToken != nil {
		return errors.New("[GetAuthToken]" + errToken.Error())
	}
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
		RequestId:    g.RequestId(false),
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

func (g *GetuiPush) RequestId(useUuid bool) (s string) {
	rand.Seed(time.Now().UnixNano())
	s = fmt.Sprintf("%s%6.0f", time.Now().Format("20190102150405"), float64(rand.Intn(999999)))
	if useUuid {
		s = uuid.NewV4().String()
	}
	return
}
