package getuigo

import "testing"
import "github.com/geek-go/getui"

func TestPushlog(t *testing.T) {
	res := getui.PushSingleResult{
		Result: "OK",
		TaskId: "12345",
		Desc:   "description",
		Status: "1",
	}

	rt := logSinglePush(&res)
	expected := "[Getui]task: 12345 status: 1 result: OK desc: description"
	if rt != expected {
		t.Log("pushlog result is different from expected")
		t.Fail()
	}
}
