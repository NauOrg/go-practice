package main

import (
	"encoding/json"
	"net/http"
	"redis-project/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	r := gin.Default()
	// ctx := services.Ctx
	rdb := services.RDB

	userApi := r.Group("/user")
	{
		userApi.GET("/", func(ctx *gin.Context) {
			iter := rdb.Scan(ctx, 0, "user:*", 0).Iterator()

			var users []User

			for iter.Next(ctx) {
				val, err := rdb.Get(ctx, iter.Val()).Result()
				if err != nil {
					continue
				}

				var user User
				json.Unmarshal([]byte(val), &user)
				users = append(users, user)
			}

			ctx.JSON(200, users)
		})

		userApi.GET("/:id", func(ctx *gin.Context) {
			var user User
			key := "user:" + ctx.Param("id")
			userRes, err := rdb.Get(ctx, key).Result()
			if err == redis.Nil {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
				return
			}
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
				return
			}
			_ = json.Unmarshal([]byte(userRes), &user)
			ctx.JSON(http.StatusOK, gin.H{"user": user})
		})
		userApi.POST("/", func(ctx *gin.Context) {
			var user User
			ctx.BindJSON(&user)
			marshalUser, _ := json.Marshal(user)
			key := "user:" + strconv.Itoa(user.Id)
			rdb.Set(ctx, key, marshalUser, time.Minute*5)
			ctx.JSON(http.StatusOK, gin.H{"mesage": "user saved", "user": user})

		})
		userApi.PUT("/", func(ctx *gin.Context) {
			var user User
			ctx.BindJSON(&user)
			key := "user:" + strconv.Itoa(user.Id)

			_, err := rdb.Get(ctx, key).Result()
			if err == redis.Nil {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
				return
			}

			marshalUser, _ := json.Marshal(user)
			rdb.Set(ctx, key, marshalUser, time.Minute*5)
			ctx.JSON(http.StatusOK, gin.H{"mesage": "user updated", "user": user})
		})
		userApi.DELETE("/:id", func(ctx *gin.Context) {
			key := "user:" + ctx.Param("id")

			if _, err := rdb.Get(ctx, key).Result(); err == redis.Nil {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
				return
			}

			res, err := rdb.Del(ctx, key).Result()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"mesage": "user deleted", "res": res})

		})
	}

	r.Run(":8080")
}
