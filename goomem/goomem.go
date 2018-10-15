package goomem

import (
	"reflect"
	"sync"
)

var incrMux = sync.RWMutex{}

//Set 设置一个值
func Set(key string, value interface{}, expire int64) bool {
	return addItem(NewMItem(key, value, expire))
}

//获取值
func Get(key string) interface{} {
	return getItemValue(key)
}

func Incr(key string) int64 {
	val := Get(key)
	incrMux.Lock()
	defer incrMux.Unlock()
	if !isBaseType(val) {
		Set(key, 0, -1)
		return 0
	}

	vof := reflect.ValueOf(val)
	//kind:=vof.Kind()
	in := vof.Int()
	in++
	Set(key, in, -1)
	return in

}
