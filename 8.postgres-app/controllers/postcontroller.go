package controllers

import (
	"net/http"
	"postgres-project/models"
	"postgres-project/services"

	"github.com/gin-gonic/gin"
)

func GetPostAll(ctx *gin.Context) {
	var posts []models.Post
	if err := services.DB.Find(&posts).Error; err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "unable to fetch"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"posts": posts})
}
func GetPostById(ctx *gin.Context) {
	id := ctx.Param("id")
	var post models.Post
	if err := services.DB.First(&post, id).Error; err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "unable to fetch"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"post": post})
}
func CreatePost(ctx *gin.Context) {
	var post models.Post
	ctx.BindJSON(&post)
	tx := services.DB.Begin()

	var user models.User
	if err := tx.First(&user, post.UserId).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "unable to verify user"})
		return
	}
	if err := tx.Save(&post).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "unable to create"})
		return
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "unable to commit"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Post": post})
}
func DeletePost(ctx *gin.Context) {
	id := ctx.Param("id")
	var post models.Post
	tx := services.DB.Begin()

	if err := tx.First(&post, id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "unable to fetch"})
		return
	}
	if err := tx.Delete(&post).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "unable to delete"})
		return
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "unable to commit"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Post": post, "message": "post deleted"})
}
