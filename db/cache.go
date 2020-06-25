package db

import (
	"sync"

	lru "github.com/hashicorp/golang-lru"
)

type Cache interface {
	Add(key, value interface{})
	Contains(key interface{}) bool
	Get(key interface{}) (interface{}, bool)
	Remove(key interface{})
}

type LRUCache struct {
	i *lru.Cache
}

func NewLRUCache(size int) Cache {
	cache, err := lru.New(size)
	if err != nil {
		panic(err)
	}
	return &LRUCache{cache}
}

func (l *LRUCache) Add(key, value interface{}) {
	_ = l.i.Add(key, value)
}

func (l *LRUCache) Contains(key interface{}) bool {
	return l.i.Contains(key)
}

func (l *LRUCache) Get(key interface{}) (interface{}, bool) {
	return l.i.Get(key)
}

func (l *LRUCache) Remove(key interface{}) {
	_ = l.i.Remove(key)
}

type MapCache struct {
	i *sync.Map
}

func NewMapCache() Cache {
	return &MapCache{&sync.Map{}}
}

func (m *MapCache) Add(key, value interface{}) {
	m.i.Store(key, value)
}

func (m *MapCache) Contains(key interface{}) bool {
	_, ok := m.i.Load(key)
	return ok
}

func (m *MapCache) Get(key interface{}) (interface{}, bool) {
	return m.i.Load(key)
}

func (m *MapCache) Remove(key interface{}) {
	m.i.Delete(key)
}
