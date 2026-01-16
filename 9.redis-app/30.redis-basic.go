package main

import (
	"fmt"
	"redis-project/services"
	"strconv"
	"sync"
	"time"
)

func main() {
	ctx, rdb := services.Ctx, services.RDB

	id := 1
	cacheKey := "user:" + strconv.Itoa(id)

	val, err := rdb.Get(ctx, cacheKey).Result()
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
	}

	fmt.Println(val)
	fmt.Println([]byte(val))

	// Fetch from DB (simulate)
	userJSON := `{"id":1,"name":"Naushad"}`

	// Store in Redis
	rdb.Set(ctx, cacheKey, userJSON, time.Minute*5)

	fmt.Println(userJSON)
	fmt.Println([]byte(userJSON))

	fmt.Println("After setting key")

	val1, err := rdb.Get(ctx, cacheKey).Result()
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
	}

	fmt.Println(val1)
	fmt.Println([]byte(val1))

	//hash
	rdb.HSet(ctx, "session:1", map[string]interface{}{
		"user_id": 1,
		"email":   "user@test.com",
	})

	fmt.Println("Hash get")
	fmt.Println(rdb.HGet(ctx, "session:1", "user_id").Result())
	fmt.Println(rdb.HGet(ctx, "session:1", "email").Result())

	//pub/sub

	var wg sync.WaitGroup
	wg.Add(1)

	sub := rdb.Subscribe(ctx, "notifications")
	ch := sub.Channel()

	go func() {
		defer wg.Done()
		for msg := range ch {
			fmt.Println("Received:", msg.Payload)
		}
	}()

	time.Sleep(time.Second)

	rdb.Publish(ctx, "notifications", "User signed up")
	rdb.Publish(ctx, "notifications", "Order placed")

	wg.Wait()
}
