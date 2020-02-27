package getuigo

import (
	"fmt"
	"github.com/geek-go/getui"
	"github.com/gogap/logs"
)

func pushlog(rt *getui.PushSingleResult) string {
	str := fmt.Sprintf("task: %s status: %s result: %s desc: %s", rt.TaskId, rt.Status, rt.Result, rt.Desc)
	defer logs.Info(str)
	return str
}
