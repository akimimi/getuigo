package getuigo

import (
	"strings"
	"testing"
)

func TestGetuiPush_RequestId(t *testing.T) {
	getui := GetuiPush{}
	rid := getui.RequestId(true)
	if !strings.Contains(rid, "-") {
		t.Errorf("request id expect uuid, actual %s", rid)
		t.Fail()
	}
	rid = getui.RequestId(false)
	if !strings.Contains(rid, "-") {
		t.Errorf("request id expect uuid, actual %s", rid)
		t.Fail()
	}
}
