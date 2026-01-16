package main

import (
	"redis-project/services"
)

func main() {
	ctx, rdb := services.Ctx, services.RDB

	rdb.Publish(ctx, "notifications", "User signed up")
	rdb.Publish(ctx, "notifications", "Order placed")

}
