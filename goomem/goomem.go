package goomem

//Set 设置一个值
func Set(key string, value interface{}, expire int64) bool {
	return addItem(NewMItem(key, value, expire))
}

//获取值
func Get(key string) interface{} {
	return getItemValue(key)
}

