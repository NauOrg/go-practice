package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/channel-stream", streamChannelHandler)

	r.Run(":8080")
}

func produceData(ch chan<- string) {
	for i := 0; i < 50; i++ {
		ch <- fmt.Sprintf("message-%d", i)
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}

func streamChannelHandler(ctx *gin.Context) {
	ch := make(chan string)

	go produceData(ch)
	for chString := range ch {
		ctx.Writer.WriteString(chString + "\n")
	}
	ctx.Writer.Flush()

}
