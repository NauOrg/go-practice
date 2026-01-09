package main

import (
	"fmt"
	"log"
	"mongo-project/routes"
	"net/http"
	"os"
	"os/signal"

	_ "mongo-project/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/hello", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "hello world!"})
		})
	}

	routes.RegisterUserRoutes(r)
	routes.RegisterPostRoutes(r)

	srvError := make(chan error, 1)
	go func() {
		srvError <- r.Run(":8080")
	}()

	//wait for Ctrl
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	select {
	case <-quit:
		fmt.Println("server shutting down!")
	case err := <-srvError:
		log.Fatalf("server error: %v", err)
	}
}

// MONGO_URI=mongodb://root:example@localhost:27017/?authSource=admin go run main.go
