package data

import (
	"context"
	"errors"
	"task_managment_api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService interface {
	GetTasks(ctx context.Context) ([]models.Task, error)
	GetTaskByID(ctx context.Context, taskID string) (models.Task, error)
	AddTask(ctx context.Context, task models.Task) error
	UpdateTaskByID(ctx context.Context, taskID string, updatedTask models.Task) error
	DeleteTaskByID(ctx context.Context, taskID string) error
}

type MongoTaskService struct {
	collection *mongo.Collection
}

func NewMongoTaskService(db *mongo.Database) TaskService {
	return &MongoTaskService{
		collection: db.Collection("tasks"),
	}
}

func (ts *MongoTaskService) GetTasks(ctx context.Context) ([]models.Task, error) {
	var tasks []models.Task
	cursor, err := ts.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (ts *MongoTaskService) GetTaskByID(ctx context.Context, taskID string) (models.Task, error) {
	var task models.Task
	objectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return models.Task{}, err
	}

	err = ts.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Task{}, errors.New("task not found")
		}
		return models.Task{}, err
	}

	return task, nil
}

func (ts *MongoTaskService) AddTask(ctx context.Context, task models.Task) error {
	_, err := ts.collection.InsertOne(ctx, task)
	return err
}

func (ts *MongoTaskService) UpdateTaskByID(ctx context.Context, taskID string, updatedTask models.Task) error {
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

	result, err := ts.collection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": update})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}

func (ts *MongoTaskService) DeleteTaskByID(ctx context.Context, taskID string) error {
	objectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return err
	}

	result, err := ts.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}
