package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

func GetClient() *mongo.Client {
	return client
}

func GetCollection(collectionName string) *mongo.Collection {
	return client.Database("academy").Collection(collectionName)
}

func Disconnect() {
	log.Println("MongoDB disconnecting...")
	defer client.Disconnect(context.Background())
}

func Connect(connectionUrl string) error {
	log.Println("MongoDB connecting...")
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(connectionUrl))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	// Ping to make sure MongoDB is connected
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}
	log.Println("MongoDB connected!")

	return nil
}
