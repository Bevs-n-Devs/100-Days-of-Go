package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func redisCaching() {
	fmt.Println("\nUsing Third-Party Libraries: Redis")

	intro := `
 Redis is an in-memory data store used for distributed caching.
	`
	fmt.Println(intro)

	rds := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
	})

	// set a key with expiration
	err := rds.Set(ctx, "greeting", "hello world", 10*time.Second).Err()
	if err != nil {
		panic(err)
	}

	// retrieve the value
	val, err := rds.Get(ctx, "greeting").Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("value:", val)
	}

	// wait fro key to expire
	time.Sleep(11 * time.Second)

	val, err = rds.Get(ctx, "greeting").Result()
	if err == redis.Nil {
		fmt.Println("Key expired")
	} else if err != nil {
		panic(err)
	}
}
