package goomem

var isUp bool = false
var memCache *gm_cache

//启动程序
func StartUp() bool {
	if isUp {
		return false
	}
	//已经启动了。
	isUp = true
	memCache = new(gm_cache)
	return true
}

//获取值
func Get(key string) (interface{}, *gm_error) {
	v, ok := memCache.Get(key)
	if !ok {
		return nil, newError(ERR_MSG_NOT_FOUND, ERR_CODE_NOT_FOUND)
	}
	if v.IsExpire() {
		memCache.Del(key)
		return nil, newError(ERR_MSG_EXPIRED, ERR_CODE_EXPIRED)
	}
	return v.Data, nil
}

//Set value
func Set(key string, value interface{}, expire int64) interface{} {
	if expire < 0 || value == nil {
		memCache.Del(key)
		return nil
	}
	memCache.Items[key] = newItem(value, expire)
	return value
}