package getuigo

import (
	"encoding/json"
	"errors"
	"github.com/geek-go/getui"
)

type GetuiConfig struct {
	AppId          string `json:"appid" yaml:"appid"`
	AppKey         string `json:"appkey" yaml:"appkey"`
	AppSecret      string `json:"appsecret" yaml:"appsecret"`
	MasterSecret   string `json:"mastersecret" yaml:"mastersecret"`
	IntentTemplate string `json:"intent_template" yaml:"intent_template"`
}

var getuiInstance *GetuiPush
var buildInstanceCnt int

func NewGeTui(config *GetuiConfig) (*GetuiPush, error) {
	if config.AppId == "" || config.AppSecret == "" || config.AppKey == "" {
		return nil, errors.New("getui config is not provided")
	}

	if getuiInstance == nil || config.AppId != getuiInstance.Config.AppId {
		getuiInstance = &GetuiPush{Config: config}
		buildInstanceCnt++
	}
	return getuiInstance, nil
}

func IGtTransmissionTemplate(payload Payload) (*getui.Transmission, *getui.PushInfo, error) {
	payloadByte, err := json.Marshal(payload)
	if err != nil {
		return nil, nil, err
	}

	// build Transmission data
	template := &getui.Transmission{
		TransmissionType:    false,
		TransmissionContent: string(payloadByte),
	}

	if payload.GetIsShowNotify() == 1 {
		// notify for multi mobile push service
		template.Notify = &getui.Notify{
			Title:   payload.GetNotifyTitle(),
			Content: payload.GetNotifyBody(),
			Intent:  payload.GetIntent(),
			Type:    NotifyTypeIntent,
		}
	}

	// config apns for ios devices
	apn := getui.Apns{Category: "ACTIONABLE"}
	if payload.GetIsShowNotify() == 1 {
		alertmsg := &getui.Alert{}
		alertmsg.Title = payload.GetNotifyTitle()
		alertmsg.Body = payload.GetNotifyBody()
		apn.Alert = alertmsg
		apn.Sound = ""
		apn.AutoBadge = "+1" //角标
		apn.ContentAvailable = 0

	} else {
		apn.Sound = "com.gexin.ios.silence"
		apn.AutoBadge = "+0" //角标
		apn.ContentAvailable = 1
	}

	pushInfo := getui.PushInfo{}
	pushInfo["aps"] = apn
	pushInfo["payload"] = string(payloadByte)

	return template, &pushInfo, nil
}
