package getuigo

import (
	"github.com/geek-go/getui"
	"time"
)

type GetuiPush struct {
	Config *GetuiConfig
}

func (this *GetuiPush) SendTransmissionByCid(cid string, payload Payload) error {

	// get auth token
	token, _ := getui.GetGeTuiToken(this.Config.AppId, this.Config.AppKey, this.Config.MasterSecret)

	// build message body
	message := getui.GetMessage()
	message.AppKey = this.Config.AppKey
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
		RequestId:    this.RequestId(),
	}

	res, err := getui.PushSingle(this.Config.AppId, token, pushSingleParam)
	if err != nil {
		return err
	}

	defer pushlog(res)
	return nil
}

func (g *GetuiPush) RequestId() string {
	return time.Now().Format("20160102150405")
}
