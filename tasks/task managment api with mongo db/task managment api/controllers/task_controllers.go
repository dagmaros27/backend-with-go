package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	services "task_managment_api/data"
	models "task_managment_api/models"
)

func GetTasks(c *gin.Context) {
	tasks, err := services.GetTasks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	if len(tasks) == 0 {
		c.JSON(http.StatusOK, gin.H{"Message": "No task is added yet"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func GetTaskById(c *gin.Context) {

	id := c.Param("id")
	task, err := services.GetTasksByTaskID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func UpdateTaskById(c *gin.Context) {
	id := c.Param("id")
	var json models.Task
	err := c.BindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = services.UpdateTaskById(id, json)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})

}

func DeleteTaskById(c *gin.Context) {

	id := c.Param("id")

	err := services.DeleteTaskById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func AddTask(c *gin.Context) {

	var json models.Task
	err := c.BindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = services.AddTask(json)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{"message": "task added successfully"})
}
