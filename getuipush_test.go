package getuigo

import (
	"strings"
	"testing"
	"time"
)

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

func TestGetuiPush_SendTransmissionByCidList(t *testing.T) {
	payload := BasicPayload{
		NotifyTitle:  "test",
		NotifyBody:   "test",
		IsShowNotify: 0,
	}

	cids := []string{"0c79f4391dc626a5480fa010777cedd2"}
	getui := GetuiPush{
		Config: &GetuiConfig{
			AppId:        "yNLX9pFEWY9hA9KFGlj2n",
			AppKey:       "eLP3QSCp3M7BMUM4WWOhj2",
			AppSecret:    "XfcSQGmtbD6xjMMcpj94f7",
			MasterSecret: "OU9YH3omlZ617Y3VEqUok1",
		}}
	if err := getui.SendTransmissionByCidList(cids, &payload); err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestGetuiPush_GetAuthToken(t *testing.T) {
	getui := GetuiPush{
		Config: &GetuiConfig{
			AppId:        "yNLX9pFEWY9hA9KFGlj2n",
			AppKey:       "eLP3QSCp3M7BMUM4WWOhj2",
			AppSecret:    "XfcSQGmtbD6xjMMcpj94f7",
			MasterSecret: "OU9YH3omlZ617Y3VEqUok1",
		}}
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
	getui := GetuiPush{
		Config: &GetuiConfig{
			AppId:        "yNLX9pFEWY9hA9KFGlj2n",
			AppKey:       "eLP3QSCp3M7BMUM4WWOhj2",
			AppSecret:    "XfcSQGmtbD6xjMMcpj94f7",
			MasterSecret: "OU9YH3omlZ617Y3VEqUok1",
		}}
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
