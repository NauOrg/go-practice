package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api := r.Group("/api")
	{

		api.GET("/hello", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"msg": "Hello world!"})
		})

		api.GET("/getId/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			ctx.JSON(200, gin.H{"ID": id})
		})
		api.GET("/search", func(ctx *gin.Context) {
			q := ctx.Query("q")
			ctx.JSON(200, gin.H{"query": q})
		})

		type User struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}

		api.POST("/user", func(ctx *gin.Context) {
			var u User
			if err := ctx.BindJSON(&u); err != nil {
				return
			}
			ctx.JSON(200, gin.H{"user": u, "name": u.Name})
		})
	}

	r.GET("/secure", authMidleware, func(ctx *gin.Context) {
		auth, _ := ctx.Get("auth")
		ctx.JSON(200, gin.H{"msg": "Welcome authorized user", "auth": auth})
	})

	r.POST("/upload", func(ctx *gin.Context) {
		file, _ := ctx.FormFile("file")
		filePath := "/Users/naushadansari/Documents/DSA/golang/5.gin/" + file.Filename
		ctx.SaveUploadedFile(file, filePath)
		ctx.JSON(200, gin.H{
			"msg":      "file uploaded successfully",
			"fileName": file.Filename,
			"path":     filePath,
		})
	})
	r.POST("/multi-upload", func(ctx *gin.Context) {
		form, _ := ctx.MultipartForm()
		files := form.File["file"]
		filePaths := map[string]string{}
		for _, file := range files {
			filePath := "/Users/naushadansari/Documents/DSA/golang/5.gin/" + file.Filename
			filePaths[file.Filename] = filePath
			ctx.SaveUploadedFile(file, filePath)
		}
		ctx.JSON(200, gin.H{
			"msg":  "files uploaded successfully",
			"path": filePaths,
		})
	})

	r.GET("/download/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		filepath := "./uploads/" + filename

		c.Header("Content-Type", "application/octet-stream")
		// c.Header("Content-Disposition", "attachment; filename="+filename)
		c.Header("Content-Transfer-Encoding", "binary")

		// c.File(filepath)
		c.FileAttachment(filepath, filename)
	})

	r.Run(":8080")
}

func authMidleware(ctx *gin.Context) {
	token := ctx.GetHeader("AUTH")
	auth := ctx.GetHeader("Authorization")
	fmt.Println(auth)
	if token != "1234" {
		ctx.JSON(401, gin.H{"error": "Unauthorize"})
		ctx.Abort()
	}
	ctx.Set("auth", strings.SplitN(auth, " ", 2))
	ctx.Next()

}
