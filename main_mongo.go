package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://0.0.0.0:27017"))
	if err != nil {
		log.Fatalln(err.Error())
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := client.Connect(ctx); err != nil {
		log.Fatalln(err.Error())
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalln(err.Error())
	}
	collection := client.Database("go-redis").Collection("customers").FindOne(ctx, bson.M{"name": "Honey"})
	var result struct {
		Name string
		Age int64
		Cars []string
	}
	err = collection.Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}