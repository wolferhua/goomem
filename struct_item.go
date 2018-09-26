package goomem

import "time"

type gm_item struct {
	Data   interface{}
	Expire int64
}

//判断数据是否过期
func (this *gm_item) IsExpire() bool {
	//永不过期
	if this.Expire == 0 {
		return false
	}
	//判断过期时间
	return this.Expire < time.Now().Unix()
}

func newItem(data interface{}, expire int64) *gm_item {
	if expire != 0 {
		expire += time.Now().Unix()
	}
	return &gm_item{
		Data:   data,
		Expire: expire,
	}
}
