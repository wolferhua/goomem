package goomem

import (
	"reflect"
	"time"
)

//数据结构
type mItem struct {
	Key      string       //标识
	Val      interface{}  //值
	Expire   int64        //过期时间
	Zone     int          //数据所在片
	TypeKind reflect.Kind //数据类型
}

//是否过期
func (item *mItem) IsExpire() bool {
	//存在过期时间，且过期时间大于当前时间
	return item.Expire > 0 && item.Expire < time.Now().Unix()
}

func NewMItem(key string, val interface{}, expire int64) *mItem {
	if expire > 0 {
		expire = time.Now().Unix() + expire
	}
	zone := 0 // crc64.Checksum([]byte(key), crc64.MakeTable(1)) % Zone
	return &mItem{
		Key:      key,
		Val:      val,
		Expire:   expire,
		Zone:     int(zone),
		TypeKind: reflect.TypeOf(val).Kind(),
	}
}
