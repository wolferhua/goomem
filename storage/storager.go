package storage

import (
	"github.com/wolferhua/goomem/item"
	"sync"
	"time"
)

//map
var storage map[string]*item.Item

//锁
var lock sync.RWMutex

//设置元素
func Set(e *item.Item) {
	Lock()
	defer Unlock()
	storage[e.Key] = e
}

//获取元素
func Get(key string) (e *item.Item, ok bool) {

	RLock()
	defer RUnlock()

	e, ok = storage[key]
	if !e.IsExpire() {
		return
	}
	return nil, false
}

//批量获取
func Gets(keys ...string) []*item.Item {
	RLock()
	defer RUnlock()

	var items []*item.Item
	if len(keys) > 0 {
		now := time.Now().Unix()
		for _, key := range keys {
			e, ok := storage[key]
			if ok && !e.IsExpireByTime(now) {
				items = append(items, e)
			}
		}
	}
	return items
}

//批量设置
func Sets(items ...*item.Item) {
	Lock()
	defer Unlock()

	//插入数据
	if len(items) > 0 {
		for i := range items {
			storage[items[i].Key] = items[i]
		}
	}
}

//锁定
func Lock() {
	lock.Lock()
}

//解锁
func Unlock() {
	lock.Unlock()
}

//锁定
func RLock() {
	lock.RLock()
}

//解锁
func RUnlock() {
	lock.RUnlock()
}
