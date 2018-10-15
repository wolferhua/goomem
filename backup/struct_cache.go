package backup

import "sync"

type gm_cache struct {
	Items map[string]gm_item
	Mu    sync.RWMutex
}

//判断key 是否存在
func (this *gm_cache) exists(key string) bool {
	_, ok := this.Items[key]
	return ok
}

func (this *gm_cache) get(key string) (v gm_item, ok bool) {
	v, ok = this.Items[key]
	return
}

func (this *gm_cache) del(key string) bool {
	delete(this.Items, key)
	return true
}

func newGmCache() *gm_cache {
	return &gm_cache{
		Mu:    *new(sync.RWMutex),
		Items: make(map[string]gm_item),
	}
}
