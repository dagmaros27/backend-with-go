package data

import (
	"errors"
	"strconv"
	models "task_managment_api/models"
)

var tasks []models.Task


func GetTasks()[]models.Task{
	return tasks
}

func GetTasksByTaskID(taskID string) (models.Task,error){

	for _,val := range tasks{
		if val.ID == taskID {
			return val,nil
		}
	}
	
	return models.Task{}, errors.New("task not found")

}


func UpdateTaskById(taskID string, updatedTask models.Task)error{
	for idx,val := range tasks{
		if val.ID == taskID {
			if updatedTask.Title != ""{
				tasks[idx].Title = updatedTask.Title
			}
			if updatedTask.Description != ""{
				tasks[idx].Description = updatedTask.Description
			}
			if updatedTask.Status != ""{
				tasks[idx].Status = updatedTask.Status
			}
			if updatedTask.DueDate != ""{
				tasks[idx].DueDate = updatedTask.DueDate
			}
			return nil
		}
	}

	return errors.New("task not found")
}


func DeleteTaskById(taskID string) error{
	for idx,val := range tasks{
		if val.ID == taskID {
			tasks = append(tasks[:idx],tasks[idx+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}


func AddTask(task models.Task) {
	task.ID = strconv.Itoa(len(tasks)+1)
	tasks = append(tasks, task)
}


