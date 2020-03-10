package getuigo

import (
	"encoding/json"
	"errors"
	"github.com/geek-go/getui"
)

type GetuiConfig struct {
	AppId        string `json:"appid" yaml:"appid"`
	AppKey       string `json:"appkey" yaml:"appkey"`
	AppSecret    string `json:"appsecret" yaml:"appsecret"`
	MasterSecret string `json:"mastersecret" yaml:"mastersecret"`
}

type BasicPayload struct {
	PushTitle    string `json:"push_title"`
	PushBody     string `json:"push_body"`
	IsShowNotify int    `json:"is_show_notify"`
	Ext          string `json:"ext"`
}

func NewGeTui(config *GetuiConfig) (*GetuiPush, error) {
	if config.AppId == "" || config.AppSecret == "" || config.AppKey == "" {
		return nil, errors.New("Getui config is not provided.")
	}

	gt := &GetuiPush{Config: config}
	return gt, nil
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

	// config apns for ios devices
	apn := getui.Apns{Category: "ACTIONABLE"}
	if payload.GetIsShowNotify() == 1 {
		alertmsg := &getui.Alert{}
		alertmsg.Title = payload.GetPushTitle()
		alertmsg.Body = payload.GetPushBody()
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
