package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data sync.Map
}

func (c *Cache) Get(key string) (interface{}, bool) {
	return c.data.Load(key)
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.data.Store(key, value)
	go func() {
		time.Sleep(ttl)
		c.data.Delete(key) // auto-remove afgter TTL
	}()
}

func inMemoryCaching() {
	fmt.Println("\n	Simple In-Memory Cache With sync.Map")

	intro := `
The sync.Map is a thread-safe map in Go, perfect for simple caching needs.
	`
	fmt.Println(intro)

	cache := Cache{}

	cache.Set("foo", "bar", 5*time.Second)

	if val, ok := cache.Get("foo"); ok {
		fmt.Println("Found:", val)
	} else {
		fmt.Println("Not found")
	}

	time.Sleep(6 * time.Second)

	if _, ok := cache.Get("foo"); !ok {
		fmt.Println("Value expired")
	}
}
