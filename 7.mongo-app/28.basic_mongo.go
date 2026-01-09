package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	uri := "mongodb://root:example@localhost:27017/?authSource=admin"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	clientOption := options.Client().ApplyURI(uri)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Fatal("connect:", err)
	}

	defer func() {
		_ = client.Disconnect(ctx)
	}()

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("ping:", err)
	}

	fmt.Println("mongo connected!")

	db := client.Database("TEST")
	userCollection := db.Collection("users")

	//index
	_ = userDropIndexIfExists(ctx, userCollection)
	if err := createIndex(ctx, userCollection); err != nil {
		log.Fatal("CreateIndex:", err)
	}

	//insert
	if res, err := userCollection.InsertOne(ctx, bson.M{"email": "email.com", "name": "first name"}); err != nil {
		log.Fatal("Insert1:", err)
	} else {
		fmt.Println("Insert1 ID", res.InsertedID)
	}
	if res, err := userCollection.InsertOne(ctx, bson.M{"email": "email@email.com", "name": "second name"}); err != nil {
		log.Fatal("Insert2:", err)
	} else {
		fmt.Println("Insert2 ID", res.InsertedID)
	}

	//findOne
	var doc bson.M
	if err := userCollection.FindOne(ctx, bson.M{"email": "email.com"}).Decode(&doc); err != nil {
		log.Fatal("FindOne:", err)
	}
	fmt.Println("fond:", doc)

	//update
	upd := bson.M{"$set": bson.M{"name": "updated name"}}
	if res, err := userCollection.UpdateOne(ctx, bson.M{"email": "email.com"}, upd); err != nil {
		log.Fatal("Update:", err)
	} else {
		fmt.Println("updated", res)
	}

	//delete
	if res, err := userCollection.DeleteOne(ctx, bson.M{"email": "email@email.com"}); err != nil {
		log.Fatal("Delete:", res)
	} else {
		fmt.Println("deleted", res)
	}

}

func createIndex(ctx context.Context, col *mongo.Collection) error {
	mod := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}
	_, err := col.Indexes().CreateOne(ctx, mod)
	return err
}

func userDropIndexIfExists(ctx context.Context, col *mongo.Collection) error {
	indexes, err := col.Indexes().List(ctx)
	if err != nil {
		return err
	}

	for indexes.Next(ctx) {
		var idx bson.M
		if err := indexes.Decode(&idx); err != nil {
			return err
		}
		fmt.Println(`idx["name"].(string)`, idx["name"].(string))
		if strings.Contains(idx["name"].(string), "email") {
			col.Indexes().DropOne(ctx, idx["name"].(string))
		}
	}
	return nil
}
