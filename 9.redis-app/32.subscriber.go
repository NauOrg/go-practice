package main

import (
	"fmt"
	"redis-project/services"
)

func main() {
	ctx, rdb := services.Ctx, services.RDB

	sub := rdb.Subscribe(ctx, "notifications")
	ch := sub.Channel()

	for msg := range ch {
		fmt.Println("Received:", msg.Payload)
	}

}
