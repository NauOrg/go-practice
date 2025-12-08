package main

import (
	"bufio"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/upload", uploadAndStreamHandler)

	r.Run(":8080")
}

func uploadAndStreamHandler(ctx *gin.Context) {
	file, _, _ := ctx.Request.FormFile("file")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		ctx.Writer.Write([]byte(line + "\n"))
		ctx.Writer.Flush()
	}

	// reader := bufio.NewReader(file)

	// for {
	// 	line, err := reader.ReadString('\n')
	// 	ctx.Writer.Write([]byte(line))
	// 	if err != nil {
	// 		break
	// 	}
	// ctx.Writer.Flush()
	// }

}
