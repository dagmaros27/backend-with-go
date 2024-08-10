// main.go
package main

import (
	"context"
	"log"
	"task_managment_api/controllers"
	"task_managment_api/data"
	"task_managment_api/router"

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

	taskService := data.NewMongoTaskService(db)
	userService := data.NewMongoUserService(db)

	taskController := controllers.NewTaskController(taskService) 
	userController := controllers.NewUserController(userService)

	r := router.SetupRouter(taskController, userController)
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