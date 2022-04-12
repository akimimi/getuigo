package getuigo

import (
	"strings"
	"testing"
	"time"
)

var testConfig = GetuiConfig{
	AppId:        "ZZtcVHjW2M7SeyKYuq9aF8",
	AppKey:       "6c2Mp4Q0bX6Nn1JojavrW3",
	AppSecret:    "AHgxvjDoX39iHsYloCzq4",
	MasterSecret: "m006ZH1cAJ6MWOZkhavHE6",
}

func TestGetuiPush_RequestId(t *testing.T) {
	getui := GetuiPush{}
	rid := getui.RequestId(true)
	if !strings.Contains(rid, "-") {
		t.Errorf("request id expect uuid, actual %s", rid)
		t.Fail()
	}
	rid = getui.RequestId(false)
	if strings.Contains(rid, "-") {
		t.Errorf("request id expect uuid, actual %s", rid)
		t.Fail()
	}
}

func TestGetuiPush_SendTransmissionByCid(t *testing.T) {
	payload := BasicPayload{
		NotifyTitle:  "test",
		NotifyBody:   "test",
		IsShowNotify: 0,
	}
	getui := GetuiPush{Config: &testConfig}
	if err := getui.SendTransmissionByCid("0c79f4391dc626a5480fa010777cedd2", &payload); err != nil {
		t.Error(err)
	}
}

func TestGetuiPush_SendTransmissionByCidWithNotify(t *testing.T) {
	payload := BasicPayload{
		NotifyTitle:  "test",
		NotifyBody:   "test",
		IsShowNotify: 1,
	}
	getui := GetuiPush{Config: &testConfig}
	if err := getui.SendTransmissionByCid("0c79f4391dc626a5480fa010777cedd2", &payload); err != nil {
		t.Error(err)
	}
}

func TestGetuiPush_SendTransmissionByCidList(t *testing.T) {
	payload := BasicPayload{
		NotifyTitle:  "test",
		NotifyBody:   "test",
		IsShowNotify: 0,
	}

	cids := []string{"0c79f4391dc626a5480fa010777cedd2"}
	getui := GetuiPush{Config: &testConfig}
	if err := getui.SendTransmissionByCidList(cids, &payload); err != nil {
		t.Error(err)
	}
}

func TestGetuiPush_SendTransmissionToAll(t *testing.T) {
	payload := BasicPayload{
		NotifyTitle:  "test",
		NotifyBody:   "test",
		IsShowNotify: 0,
	}

	getui := GetuiPush{Config: &testConfig}
	if err := getui.SendTransmissionToAll(&payload); err != nil {
		t.Error(err)
	}
}

func TestGetuiPush_GetAuthToken(t *testing.T) {
	getui := GetuiPush{Config: &testConfig}
	token, err := getui.GetAuthToken()
	if getui.RequestTokenCnt != 1 {
		t.Errorf("GetAuthToken should request 1 time, but actual %d times", getui.RequestTokenCnt)
	}
	if err != nil || token == "" {
		t.Errorf("GetAuthToken failed (%s) with: %s", token, err.Error())
	}

	token, err = getui.GetAuthToken()
	if getui.RequestTokenCnt != 1 {
		t.Errorf("GetAuthToken should use cache, but actual request %d times", getui.RequestTokenCnt)
	}
	if err != nil || token == "" {
		t.Errorf("GetAuthToken failed (%s) with: %s", token, err.Error())
	}
}

func TestGetuiPush_IsAuthTokenValid(t *testing.T) {
	getui := GetuiPush{Config: &testConfig}
	valid := getui.IsAuthTokenValid()
	if valid != false {
		t.Error("Empty token should be invalid.")
	}
	getui.Token = "abcdefg"
	getui.TokenExpires = time.Now().Add(-1 * time.Second)
	valid = getui.IsAuthTokenValid()
	if valid != false {
		t.Error("Token should be expired.")
	}

	getui.TokenExpires = time.Now().Add(3 * time.Second)
	valid = getui.IsAuthTokenValid()
	if valid != true {
		t.Error("Token should be valid.")
	}
}

func TestNewGeTui(t *testing.T) {
	instance, err := NewGeTui(&testConfig)
	if instance == nil || err != nil {
		t.Errorf("NewGetui failed, %v, error: %s", instance, err.Error())
	}
	if getuiInstance == nil {
		t.Error("Static variable should not be nil.")
	}
	instance2, err := NewGeTui(&testConfig)
	if instance2 != instance || buildInstanceCnt > 1 {
		t.Error("Instance should not be built again.")
	}
}
