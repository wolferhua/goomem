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
func (i Item) IsExpire() bool {
	if i.Expire > 0 && i.Expire < time.Now().Unix() {
		return true
	}
	return false
}
