// main.go
package main

import (
	"context"
	"log"
	"task_managment_api/delivery/router"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("task_manager")
	EnsureIndexes(db)
	

	r := router.SetupRouter(db, time.Second * 3)
	r.Run(":8080")
}



//to enforce username uniqueness in the db
func EnsureIndexes(db *mongo.Database) error {
	userCollection := db.Collection("users")
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"username": 1}, 
		Options: options.Index().SetUnique(true),
	}

	_, err := userCollection.Indexes().CreateOne(context.TODO(), indexModel)
	return err
}