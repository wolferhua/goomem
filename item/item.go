package item

import "time"

// 数据类型
type Type int

const (
	_ Type = iota
	TypeString
	TypeList
	TypeHash
)

type Item struct {
	Key    string
	Value  interface{}
	Type   Type
	Expire int64
}

//判断是否过期
func (e *Item) IsExpire() bool {
	if e.Expire > 0 && e.Expire < time.Now().Unix() {
		return true
	}
	return false
}

//验证过期
func (e *Item) IsExpireByTime(time int64) bool {
	if e.Expire > 0 && e.Expire < time {
		return true
	}
	return false
}
