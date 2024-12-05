package main

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

func goCache() {
	fmt.Println("\n	Using Third Party Libraries: GoCache")

	intro := `
GoCache (github.com/patrickmn/go-cache): A lightweight in-memory cache with expiration handling.	
	`
	fmt.Println(intro)

	c := cache.New(5*time.Minute, 10*time.Minute)

	c.Set("username", "john_doe", cache.DefaultExpiration)

	value, found := c.Get("username")
	if found {
		fmt.Println("cached value:", value)
	}

	time.Sleep(6 * time.Minute)

	_, found = c.Get("username")
	if !found {
		fmt.Println("Cache expired")
	}
}
