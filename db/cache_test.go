package db_test

import (
	"testing"

	"github.com/go-snart/snart/db"
)

func TestNewLRUCachePanic(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Fatal("err shouldn't be nil")
		}
	}()

	_ = db.NewLRUCache(-1)
}

func TestLRUCacheGet(t *testing.T) {
	c := db.NewLRUCache(1)

	c.Set("a", 1)

	value := c.Get("a")
	if value != 1 {
		t.Fatal("value isn't 1")
	}
}

func TestLRUCacheHas(t *testing.T) {
	c := db.NewLRUCache(1)

	c.Set("a", 1)

	if !c.Has("a") {
		t.Fatal("cache should contain a")
	}
}

func TestLRUCacheDel(t *testing.T) {
	c := db.NewLRUCache(1)

	c.Set("a", 1)
	c.Del("a")

	if c.Has("a") {
		t.Fatal("cache shouldn't contain a")
	}
}

func TestLRUCacheKeys(t *testing.T) {
	c := db.NewLRUCache(1)

	c.Set("a", 1)

	keys := c.Keys()

	if len(keys) != 1 {
		t.Fatalf("len(keys) == %d != 1", len(keys))
	}

	if keys[0] != "a" {
		t.Fatalf("keys[0] == %#v != a", keys[0])
	}
}

func TestMapCacheGet(t *testing.T) {
	c := db.NewMapCache()

	c.Set("a", 1)

	value := c.Get("a")
	if value != 1 {
		t.Fatal("value isn't 1")
	}
}

func TestMapCacheDel(t *testing.T) {
	c := db.NewMapCache()

	c.Set("a", 1)
	c.Del("a")

	if c.Has("a") {
		t.Fatal("cache shouldn't contain a")
	}
}

func TestMapCacheKeys(t *testing.T) {
	c := db.NewMapCache()
	c.Set("a", 1)

	keys := c.Keys()

	if len(keys) != 1 {
		t.Fatalf("len(keys) == %d != 1", len(keys))
	}

	if keys[0] != "a" {
		t.Fatalf("keys[0] == %#v != a", keys[0])
	}
}
