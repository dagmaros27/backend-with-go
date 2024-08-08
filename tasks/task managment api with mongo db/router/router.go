package router

import (
	"task_managment_api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(taskController *controllers.TaskController) *gin.Engine {
	router := gin.Default()

	router.GET("/tasks", taskController.GetTasks)
	router.GET("/tasks/:id", taskController.GetTaskByID)
	router.POST("/tasks", taskController.AddTask)
	router.PUT("/tasks/:id", taskController.UpdateTaskByID)
	router.DELETE("/tasks/:id", taskController.DeleteTaskByID)

	return router
}
