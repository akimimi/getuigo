package getuigo

import (
	"strings"
	"testing"
)

func TestGetuiPush_RequestId(t *testing.T) {
	getui := GetuiPush{}
	rid := getui.RequestId()
	if !strings.Contains(rid, "-") {
		t.Errorf("request id expect uuid, actual %s", rid)
		t.Fail()
	}
}
