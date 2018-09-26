package goomem

import "time"

type item struct {
	Data   interface{}
	Expire int64
}

//判断数据是否过期
func (this *item) IsExpire() bool {
	//永不过期
	if this.Expire == 0 {
		return false
	}
	//判断过期时间
	return this.Expire < time.Now().Unix()
}
