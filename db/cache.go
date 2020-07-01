package db

import (
	"sync"

	"github.com/hashicorp/golang-lru/simplelru"
)

// Cache is a generalization of a key:value cache.
type Cache interface {
	sync.Locker

	Get(key interface{}) interface{}
	Set(key, value interface{})
	Del(key interface{})
	Has(key interface{}) bool
	Keys() []interface{}
}

// LRUCache is a Cache using github.com/hashicorp/golang-lru/simplelru.
type LRUCache struct {
	*sync.Mutex

	i *simplelru.LRU
}

var _ Cache = (*LRUCache)(nil)

// NewLRUCache creates a *LRUCache with a given size.
func NewLRUCache(size int) *LRUCache {
	cache, err := simplelru.NewLRU(size, nil)
	if err != nil {
		panic(err)
	}
	return &LRUCache{&sync.Mutex{}, cache}
}

// Get retrieves a value from a given key.
func (l *LRUCache) Get(key interface{}) interface{} {
	value, _ := l.i.Get(key)
	return value
}

// Set assigns a value to a given key.
func (l *LRUCache) Set(key, value interface{}) {
	_ = l.i.Add(key, value)
}

// Del removes a value from a given key.
func (l *LRUCache) Del(key interface{}) {
	_ = l.i.Remove(key)
}

// Has reports if a given key is present.
func (l *LRUCache) Has(key interface{}) bool {
	return l.i.Contains(key)
}

// Keys returns a slice of all keys present.
func (l *LRUCache) Keys() []interface{} {
	return l.i.Keys()
}

// MapCache is a Cache using a map[interface{}]interface{}
type MapCache struct {
	*sync.Mutex

	i map[interface{}]interface{}
}

var _ Cache = (*MapCache)(nil)

// NewMapCache creates a *MapCache.
func NewMapCache() *MapCache {
	return &MapCache{&sync.Mutex{}, map[interface{}]interface{}{}}
}

// Get retrieves a value from a given key.
func (m *MapCache) Get(key interface{}) interface{} {
	value, _ := m.i[key]
	return value
}

// Set assigns a value to a given key.
func (m *MapCache) Set(key, value interface{}) {
	m.i[key] = value
}

// Del removes a value from a given key.
func (m *MapCache) Del(key interface{}) {
	delete(m.i, key)
}

// Has reports if a given key is present.
func (m *MapCache) Has(key interface{}) bool {
	_, ok := m.i[key]
	return ok
}

// Keys returns a slice of all keys present.
func (m *MapCache) Keys() []interface{} {
	ks, i := make([]interface{}, len(m.i)), 0
	for k := range m.i {
		ks[i], i = k, i+1
	}
	return ks
}
