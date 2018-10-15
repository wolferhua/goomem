package goomem

import (
	"fmt"
	"sync"
	"time"
)

var items = map[string]*mItem{}
var itemMux = &sync.RWMutex{}
var tk *time.Ticker

func init() {
	//设置定时器，定时过期
	tk = time.NewTicker(time.Second)
	go ticker()
}

func ticker() {
	for range tk.C {
		for key, item := range items {
			fmt.Println(key)
			//fmt.Println(item)
			if item.IsExpire() {
				itemMux.Lock()
				delete(items, key)
				itemMux.Unlock()
			}
		}
	}
}

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
	if ok && !item.IsExpire() {
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
