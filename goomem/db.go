package goomem

import "sync"

var items = map[string]*mItem{}
var itemMux = &sync.RWMutex{}

// item 加入map
func addItem(item *mItem) bool {
	if len(item.Key) == 0 {
		return false
	}
	//防止多次操作
	itemMux.Lock()
	defer itemMux.Unlock()

	//写入map
	items[item.Key] = item
	return true
}

//get item
func getItem(key string) *mItem {
	item, ok := items[key]
	if ok {
		return item
	}
	return nil
}

//获取 item 值
func getItemValue(key string) interface{} {
	item := getItem(key)
	if item == nil {
		return nil
	}
	return item.Val
}
