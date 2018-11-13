package goomem

import (
	"github.com/wolferhua/goomem/item"
	"github.com/wolferhua/goomem/storage"
	"strings"
)

func Dispatcher(cmd *Command) *Result {
	ret := &Result{}
	switch strings.ToUpper(cmd.Cmd) {
	//获取元素
	case CmdGet:
		e, ok := storage.Get(cmd.GetArg(0).(string))
		ret.Add(e).Add(ok)

	//设置元素
	case CmdSet:
		storage.Set(cmd.GetArg(0).(*item.Item))
	}
	return ret
}
