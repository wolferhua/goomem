package goomem

import "github.com/wolferhua/goomem/item"

//获取元素
func Get(key string) interface{} {
	ret := Dispatcher(&Command{
		Cmd:  CmdGet,
		Args: []interface{}{key},
	})

	if !ret.Has() {
		return nil
	}

	if ret.Get(1).(bool) {
		return nil
	}

	return ret.Get(0).(*item.Item).Value
}

//设置元素
func Set(key string, value interface{}, expire ...int64) {
	var exp int64
	if len(expire) > 0 {
		exp = expire[0]
	}
	e := item.NewItem(key, value, exp)
	Dispatcher(&Command{
		Cmd:  CmdSet,
		Args: []interface{}{e},
	})
}
