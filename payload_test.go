package getuigo

import "testing"

func TestBasicPayload_SetNotifyTitle(t *testing.T) {
	payload := BasicPayload{}
	payload.SetNotifyTitle("unittest")

	if payload.GetNotifyTitle() != "unittest" {
		t.Errorf("title expects to be %s, actual %s", "unittest", payload.GetNotifyTitle())
	}

	if payload.String() != "unittest" {
		t.Errorf("payload string value expects to be %s, actual %s", "unittest", payload.String())
	}
}

func TestBasicPayload_SetNotifyBody(t *testing.T) {
	payload := BasicPayload{}
	payload.SetNotifyBody("unittest")
	if payload.GetNotifyBody() != "unittest" {
		t.Errorf("body expects to be %s, actual %s", "unittest", payload.GetNotifyTitle())
	}
}

func TestBasicPayload_SetIsShowNotify(t *testing.T) {
	payload := BasicPayload{}
	payload.SetIsShowNotify(1)
	if payload.GetIsShowNotify() != 1 {
		t.Errorf("IsShowNotify expects to be %d, actual %d", 1, payload.GetIsShowNotify())
	}
}

func TestBasicPayload_SetIntent(t *testing.T) {
	payload := BasicPayload{}
	payload.SetIntent("unittest")
	if payload.GetIntent() != "" {
		t.Error("BasicPayload should to nothing in SetIntent")
	}
}

func TestBasicPayload_SetExt(t *testing.T) {
	payload := BasicPayload{}
	payload.SetExt("unittest")
	if payload.GetExt() != "unittest" {
		t.Errorf("ext expects to be %s, actual %s", "unittest", payload.GetExt())
	}
}
