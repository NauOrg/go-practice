package main

import (
	_ "postgres-project/models"
	_ "postgres-project/services"

	"fmt"
	"os"
	"os/signal"
	"postgres-project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.RegisterPost(r)
	routes.RegisterUser(r)

	srvError := make(chan error, 1)
	go func() {
		srvError <- r.Run(":8080")
	}()

	//wait for ctrl
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	select {
	case <-quit:
		fmt.Println("server shutting down!")
	case err := <-srvError:
		fmt.Println("server error %v", err)
	}
}
