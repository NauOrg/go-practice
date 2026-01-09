package controllers

import (
	"fmt"
	"net/http"

	"mongo-project/models"
	"mongo-project/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var Posts = []models.Post{
	{ID: primitive.NewObjectID(), Title: "test title1", Content: "test content1", UserID: primitive.NewObjectID()},
	{ID: primitive.NewObjectID(), Title: "test title2", Content: "test content2", UserID: primitive.NewObjectID()},
}
var postCollection = services.Client.Database("TEST").Collection("posts")

func GetPosts(ctx *gin.Context) {
	context := ctx.Request.Context()
	cur, err := postCollection.Find(context, bson.M{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch posts",
		})
		return
	}
	var posts = []models.Post{}
	for cur.Next(context) {
		var post models.Post
		if err := cur.Decode(&post); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to decode post",
			})
			return
		}
		posts = append(posts, post)
	}
	ctx.JSON(http.StatusOK, gin.H{"Posts": posts})
}
func GetPostById(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	context := ctx.Request.Context()
	var post = models.Post{}
	if err := postCollection.FindOne(context, bson.M{"_id": id}).Decode(&post); err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("%s id not found", id.Hex()),
			})
			return

		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch posts",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"post": post})

}

func CreatePost(ctx *gin.Context) {
	context := ctx.Request.Context()
	var post struct {
		ID      string `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
		UserID  string `json:"userId"`
	}
	ctx.BindJSON(&post)
	//verify userid
	uid, err := primitive.ObjectIDFromHex(post.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	if err := UserCollection.FindOne(context, bson.M{"_id": uid}).Decode(nil); err == mongo.ErrNoDocuments {
		ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("userId %s not found", uid.Hex())})
		return
	}

	p := &models.Post{Title: post.Title, Content: post.Content, UserID: uid}
	if res, err := postCollection.InsertOne(context, p); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create post",
			"error":   err,
		})
		return
	} else {
		p.ID = res.InsertedID.(primitive.ObjectID)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "post created successfully", "Post": p})
	}
}

func DeletePost(ctx *gin.Context) {
	context := ctx.Request.Context()
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	if res, err := postCollection.DeleteOne(context, bson.M{"_id": id}); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to delete post",
			"error":   err,
		})
		return
	} else {

		ctx.JSON(http.StatusOK, gin.H{"message": "post deleted", "Deleted": res})
		return
	}

}
