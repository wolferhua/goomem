package config

import (
	"reflect"
	"sync"
)

var initLock sync.Mutex
var initStatus = false
var conf = NewConfig()

func Init(config map[string]interface{}) {
	initLock.Lock()
	defer initLock.Unlock()
	if initStatus {
		return
	}
	t := reflect.TypeOf(conf)
	num := t.NumField()
	//config reflect
	confRef := reflect.ValueOf(conf)
	for i := 0; i < num; i++ {
		k := t.Field(i).Tag.Get("conf")
		if v, ok := config[k]; ok {
			confRef.Field(i).Set(reflect.ValueOf(v))
		}
	}
}
