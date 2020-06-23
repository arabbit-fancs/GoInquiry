package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var collection *mongo.Collection

//Mongo接続
func initMongo() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:root@localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("gotest").Collection("inquiry")
}

//書き込み
func writeInquiry(inquiry Inquiry) error {
	_, err := collection.InsertOne(context.Background(), bson.M{
		"name": inquiry.Name,
		"mail": inquiry.Mail,
		"msg":  inquiry.Msg,
	})
	return err
}
