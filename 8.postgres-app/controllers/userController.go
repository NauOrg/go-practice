package controllers

import (
	"net/http"
	"postgres-project/models"
	"postgres-project/services"

	"github.com/gin-gonic/gin"
)

func GetUserAll(ctx *gin.Context) {
	var users []models.User
	if err := services.DB.Preload("Post").Find(&users).Error; err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "unable to fetch"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}
func GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User
	if err := services.DB.Preload("Post").First(&user, id).Error; err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "unable to fetch"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
func CreateUser(ctx *gin.Context) {
	var user models.User
	ctx.BindJSON(&user)
	tx := services.DB.Begin()
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "unable to create"})
		return
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "unable to commit"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"User": user})
}

func UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var userBody models.User
	// ctx.BindJSON(&userBody)
	// var user models.User
	tx := services.DB.Begin()

	if err := tx.First(&userBody, id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "unable to fetch"})
		return
	}
	ctx.BindJSON(&userBody)
	if err := tx.Save(&userBody).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "unable to update"})
		return
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "unable to commit"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"User": userBody, "message": "user upadted"})
}

func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User
	tx := services.DB.Begin()

	if err := tx.First(&user, id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "unable to fetch"})
		return
	}
	if err := tx.Delete(&user).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "unable to delete"})
		return
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "unable to commit"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"User": user, "message": "user deleted"})
}
