package main

import (
	"context"
	"fmt"
	redis "github.com/go-redis/redis/v9"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	if redisClient == nil {
		panic("failed to nit redis client")
	}
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	key := "blog.title"

	err = redisClient.Set(ctx, key, "Use Redis to cache data in Go", 0).Err()
	if err != nil {
		panic(err)
	}

	value, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s -> %s\n", key, value)
}
