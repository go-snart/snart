package db

import lru "github.com/hashicorp/golang-lru"

type Cache interface {
	Add(key, value interface{})
	Contains(key interface{}) bool
	Get(key interface{}) (interface{}, bool)
	Remove(key interface{})
}

type LRUCache struct {
	*lru.Cache
}

func (l *LRUCache) Add(key, value interface{}) {
	_ = l.Cache.Add(key, value)
}

func (l *LRUCache) Contains(key interface{}) bool {
	return l.Cache.Contains(key)
}

func (l *LRUCache) Get(key interface{}) (interface{}, bool) {
	return l.Cache.Get(key)
}

func (l *LRUCache) Remove(key interface{}) {
	_ = l.Cache.Remove(key)
}

func NewLRUCache(size int) Cache {
	cache, err := lru.New(size)
	if err != nil {
		panic(err)
	}
	return &LRUCache{cache}
}

type MapCache map[interface{}]interface{}

func (m *MapCache) Add(key, value interface{}) {
	(*m)[key] = value
}

func (m *MapCache) Contains(key interface{}) bool {
	_, ok := (*m)[key]
	return ok
}

func (m *MapCache) Get(key interface{}) (interface{}, bool) {
	value, ok := (*m)[key]
	return value, ok
}

func (m *MapCache) Remove(key interface{}) {
	delete(*m, key)
}

func NewMapCache() Cache {
	return &MapCache{}
}
