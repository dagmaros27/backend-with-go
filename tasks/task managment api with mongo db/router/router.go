package router

import (
	"github.com/gin-gonic/gin"
	c "task_managment_api/controllers"
)


func Init(){
	router := gin.Default()
	

	//routes
	router.GET("/tasks",c.GetTasks)
	router.GET("/tasks/:id", c.GetTaskById)
	router.POST("/tasks", c.AddTask)
	router.DELETE("/tasks/:id", c.DeleteTaskById)
	router.PUT("/tasks/:id", c.UpdateTaskById)
	
	router.Run("localhost:3000")
}