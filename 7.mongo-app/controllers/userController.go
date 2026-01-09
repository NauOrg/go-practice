package controllers

import (
	"fmt"
	"mongo-project/models"
	"mongo-project/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var Users = []models.User{
	{ID: primitive.NewObjectID(), Name: "name1", Email: "email1"},
	{ID: primitive.NewObjectID(), Name: "name2", Email: "email2"},
}

var UserCollection = services.Client.Database("TEST").Collection("users")

func GetUsers(ctx *gin.Context) {
	var users []models.User
	context := ctx.Request.Context()

	cur, err := UserCollection.Find(context, bson.M{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch users",
		})
		return
	}
	defer cur.Close(context)

	for cur.Next(context) {
		var user models.User
		if err := cur.Decode(&user); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to decode user",
			})
			return
		}
		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "cursor error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func GetUserById(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	context := ctx.Request.Context()
	var user models.User
	if err := UserCollection.FindOne(context, bson.M{"_id": id}).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("%s not found", id.Hex()),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to fetch users",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"users": user,
	})
}
func CreateUser(ctx *gin.Context) {
	user := models.User{}
	ctx.BindJSON(&user)
	context := ctx.Request.Context()
	if res, err := UserCollection.InsertOne(context, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create users",
			"error":   err,
		})
		return
	} else {

		user.ID = res.InsertedID.(primitive.ObjectID)
		ctx.JSON(http.StatusOK, gin.H{"message": "user created", "User": user, "InsertedId": res})
		return
	}
}

func DeleteUser(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return

	}
	context := ctx.Request.Context()
	if res, err := UserCollection.DeleteOne(context, bson.M{"_id": id}); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to delete users",
			"error":   err,
		})
		return
	} else {

		ctx.JSON(http.StatusOK, gin.H{"message": "user deleted", "Deleted": res})
		return
	}

}
