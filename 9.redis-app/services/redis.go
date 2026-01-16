package services

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	RDB *redis.Client
	Ctx = context.Background()
)

func connectRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err := RDB.Ping(Ctx).Result()
	if err != nil {

		log.Fatal("Redis connection failed:", err)
	}
	fmt.Println("Redis connected!")
}

func init() {
	connectRedis()
}
