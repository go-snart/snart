package db

import "testing"

func TestNewLRUCachePanic(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Fatal("err shouldn't be nil")
		}
	}()

	_ = NewLRUCache(-1)
}

func TestLRUCacheGet(t *testing.T) {
	c := NewLRUCache(1)
	c.Add("a", 1)
	value, ok := c.Get("a")
	if !ok {
		t.Fatal("not ok")
	}
	if value != 1 {
		t.Fatal("value isn't 1")
	}
}

func TestLRUCacheContains(t *testing.T) {
	c := NewLRUCache(1)
	c.Add("a", 1)
	if !c.Contains("a") {
		t.Fatal("cache should contain a")
	}
}

func TestLRUCacheRemove(t *testing.T) {
	c := NewLRUCache(1)
	c.Add("a", 1)
	c.Remove("a")
	if c.Contains("a") {
		t.Fatal("cache shouldn't contain a")
	}
}

func TestMapCacheGet(t *testing.T) {
	c := NewMapCache()
	c.Add("a", 1)
	value, ok := c.Get("a")
	if !ok {
		t.Fatal("not ok")
	}
	if value != 1 {
		t.Fatal("value isn't 1")
	}
}

func TestMapCacheRemove(t *testing.T) {
	c := NewMapCache()
	c.Add("a", 1)
	c.Remove("a")
	if c.Contains("a") {
		t.Fatal("cache shouldn't contain a")
	}
}
