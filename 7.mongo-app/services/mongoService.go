package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Db *mongo.Database

func init() {
	mongoUri := os.Getenv("MONGO_URI")
	if mongoUri == "" {
		mongoUri = "mongodb://root:example@localhost:27017/?authSource=admin"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOption := options.Client().ApplyURI(mongoUri)

	var err error
	fmt.Println("Client", Client)
	Client, err = mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Fatal("Mongo Connect:", err)
	}

	if err := Client.Ping(ctx, nil); err != nil {
		log.Fatal("ping:", err)
	}

	fmt.Println("mongo connected!")

	// Db := Client.Database("TEST")
	// Db.Collection("USER")

}
