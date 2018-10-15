package goomem

import "time"

type Item struct {
	Key    string //标识
	Val    string //值
	Expire int64  //过期时间
	Zone   int    //数据所在片
}

//是否过期
func (this *Item) IsExpire() bool {
	//存在过期时间，且过期时间大于当前时间
	return this.Expire != 0 && this.Expire > time.Now().Unix()
}
