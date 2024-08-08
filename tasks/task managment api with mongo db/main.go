package main

import (
	"task_managment_api/controllers"
	"task_managment_api/data"
	"task_managment_api/router"
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	db := client.Database("task_manager")

	taskService := data.NewMongoTaskService(db)

	taskController := controllers.NewTaskController(taskService)

	r := router.SetupRouter(taskController)
	r.Run(":8080")
}
