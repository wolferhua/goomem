package goomem

import "reflect"

var isUp bool = false
var memCache *gm_cache

//启动程序
func StartUp() bool {
	if isUp {
		return false
	}
	//已经启动了。
	isUp = true
	memCache = newGmCache()
	return true
}

//获取值
func Get(key string) (interface{}, *gm_error) {
	v, ok := memCache.get(key)
	if !ok {
		return nil, newError(ERR_MSG_NOT_FOUND, ERR_CODE_NOT_FOUND)
	}
	if v.IsExpire() {
		memCache.del(key)
		return nil, newError(ERR_MSG_EXPIRED, ERR_CODE_EXPIRED)
	}
	return v.Data, nil
}

func GetForType(key string, t interface{}) *gm_error {
	v, err := Get(key)
	if err != nil {
		return err
	}
	//reflect.TypeOf(v).ConvertibleTo(t)
	t = reflect.ValueOf(v)
	return nil
}

//Set value
func Set(key string, value interface{}, expire int64) interface{} {
	memCache.Mu.Lock()
	defer memCache.Mu.Unlock()
	if expire < 0 || value == nil {
		memCache.del(key)
		return nil
	}
	memCache.Items[key] = *newItem(value, expire)
	return value
}
