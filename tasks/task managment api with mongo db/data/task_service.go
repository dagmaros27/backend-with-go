package data

import (
	"context"
	"errors"
	"log"
	//"strconv"
	models "task_managment_api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func dbInit() *mongo.Collection {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	taskCollection := client.Database("task_manager").Collection("tasks")

	return taskCollection
}

var taskCollection mongo.Collection = *dbInit()

func GetTasks() ([]models.Task, error) {
	var tasks []models.Task

	findResult, err := taskCollection.Find(context.TODO(), bson.D{})

	if err != nil {
		return nil, err
	}

	err = findResult.All(context.TODO(), &tasks)

	if err != nil {
		return nil, err
	}

	return tasks, nil

}

func GetTasksByTaskID(taskID string) (models.Task, error) {

	var task models.Task

	objectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return models.Task{}, err
	}

	err = taskCollection.FindOne(context.TODO(), bson.D{{"_id", objectID}}).Decode(&task)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Task{}, errors.New("task not found")
		}
		return models.Task{}, err
	}

	return task, nil
}


func UpdateTaskById(taskID string, updatedTask models.Task) error {
	objectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return err
	}

	update := bson.M{}
	if updatedTask.Title != "" {
		update["title"] = updatedTask.Title
	}
	if updatedTask.Description != "" {
		update["description"] = updatedTask.Description
	}
	if updatedTask.Status != "" {
		update["status"] = updatedTask.Status
	}
	if updatedTask.DueDate != "" {
		update["due_date"] = updatedTask.DueDate
	}

	if len(update) == 0 {
		return errors.New("no fields to update")
	}

	result, err := taskCollection.UpdateOne(context.TODO(), bson.M{"_id": objectID}, bson.M{"$set": update})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}


func DeleteTaskById(taskID string) error {
	objectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return err
	}

	deletedResult,err := taskCollection.DeleteOne(context.TODO(),bson.M{"_id":objectID})

	if err != nil {
		return err
	}

	if deletedResult.DeletedCount == 0{
		return errors.New("task not found")
	}

	return nil
}

func AddTask(task models.Task) error{
	_,err := taskCollection.InsertOne(context.TODO(),task)
	if err != nil {
		return err
	}
	return nil
}
