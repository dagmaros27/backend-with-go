package controllers

import (
	//"context"
	"net/http"
	"task_managment_api/data"
	"task_managment_api/models"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskService data.TaskService
}

func NewTaskController(taskService data.TaskService) *TaskController {
	return &TaskController{
		taskService: taskService,
	}
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks, err := tc.taskService.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(tasks) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No tasks found"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := tc.taskService.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) UpdateTaskByID(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := tc.taskService.UpdateTaskByID(id, task)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func (tc *TaskController) DeleteTaskByID(c *gin.Context) {
	id := c.Param("id")
	err := tc.taskService.DeleteTaskByID( id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func (tc *TaskController) AddTask(c *gin.Context) {
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := tc.taskService.AddTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Task added successfully"})
}
