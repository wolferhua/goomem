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

func NewItem(key string, value interface{}, expire int64) *Item {
	if expire > 0 {
		expire += time.Now().Unix()
	} else {
		expire = 0
	}
	return &Item{
		Key:    key,
		Value:  value,
		Expire: expire,
	}
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
