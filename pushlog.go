package getuigo

import (
	"fmt"
	"github.com/geek-go/getui"
	"github.com/gogap/logs"
)

const (
	LogNone = iota
	LogStd
	LogFile
)

var LogType = LogNone

func logSinglePush(rt *getui.PushSingleResult) string {
	str := fmt.Sprintf("task: %s status: %s result: %s desc: %s", rt.TaskId, rt.Status, rt.Result, rt.Desc)
	if LogType == LogStd {
		defer logs.Info(str)
	}
	return str
}

func logListPush(rt *getui.PushListResult) string {
	str := fmt.Sprintf("task: %s result: %s desc: %s", rt.Taskid, rt.Result, rt.Desc)
	if LogType == LogStd {
		defer logs.Info(str)
	}
	return str
}

func logAppPush(rt *getui.PushAppResult) string {
	str := fmt.Sprintf("task: %s result: %s desc: %s", rt.Taskid, rt.Result, rt.Desc)
	if LogType == LogStd {
		defer logs.Info(str)
	}
	return str
}
