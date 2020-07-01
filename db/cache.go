package db

import (
	"sync"

	"github.com/hashicorp/golang-lru/simplelru"
)

type Cache interface {
	sync.Locker

	Get(key interface{}) interface{}
	Set(key, value interface{})
	Del(key interface{})
	Has(key interface{}) bool
	Keys() []interface{}
}

type LRUCache struct {
	*sync.Mutex

	i *simplelru.LRU
}

var _ Cache = (*LRUCache)(nil)

func NewLRUCache(size int) *LRUCache {
	cache, err := simplelru.NewLRU(size, nil)
	if err != nil {
		panic(err)
	}
	return &LRUCache{&sync.Mutex{}, cache}
}

func (l *LRUCache) Get(key interface{}) interface{} {
	value, _ := l.i.Get(key)
	return value
}

func (l *LRUCache) Set(key, value interface{}) {
	_ = l.i.Add(key, value)
}

func (l *LRUCache) Del(key interface{}) {
	_ = l.i.Remove(key)
}

func (l *LRUCache) Has(key interface{}) bool {
	return l.i.Contains(key)
}

func (l *LRUCache) Keys() []interface{} {
	return l.i.Keys()
}

type MapCache struct {
	*sync.Mutex

	i map[interface{}]interface{}
}

var _ Cache = (*MapCache)(nil)

func NewMapCache() *MapCache {
	return &MapCache{&sync.Mutex{}, map[interface{}]interface{}{}}
}

func (m *MapCache) Get(key interface{}) interface{} {
	value, _ := m.i[key]
	return value
}

func (m *MapCache) Set(key, value interface{}) {
	m.i[key] = value
}

func (m *MapCache) Del(key interface{}) {
	delete(m.i, key)
}

func (m *MapCache) Has(key interface{}) bool {
	_, ok := m.i[key]
	return ok
}

func (m *MapCache) Keys() []interface{} {
	ks, i := make([]interface{}, len(m.i)), 0
	for k := range m.i {
		ks[i], i = k, i+1
	}
	return ks
}
