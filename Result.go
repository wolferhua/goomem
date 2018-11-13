package goomem

import "sync"

type Result struct {
	Total   int
	Results []interface{}
	lock    sync.RWMutex
}

func (ret *Result) Add(e interface{}) *Result {
	ret.lock.Lock()
	defer ret.lock.Unlock()
	ret.Total++
	ret.Results = append(ret.Results, e)
	return ret
}

func (ret *Result) Get(i int) interface{} {
	ret.lock.RLock()
	defer ret.lock.RUnlock()
	if i >= 0 && i < ret.Total {
		return ret.Results[i]
	}
	return nil
}

func (ret *Result) Has() bool {
	return ret.Total > 0
}
