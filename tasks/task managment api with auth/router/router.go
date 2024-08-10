package router

import (
	"task_managment_api/controllers"
	"github.com/gin-gonic/gin"
	"task_managment_api/middlewares"
)

func SetupRouter(taskController *controllers.TaskController, userController *controllers.UserController) *gin.Engine {
	router := gin.Default()

	// routes that don't require authentication
	router.POST("/register", userController.RegisterUser)
	router.POST("/login", userController.LoginUser)

	// routes that require authentication
	authorized := router.Group("/")
	authorized.Use(middlewares.AuthMiddleware())

	// task routes
	authorized.GET("/tasks", taskController.GetTasks)
	authorized.GET("/tasks/:id", taskController.GetTaskByID)
	authorized.POST("/tasks", middlewares.AdminMiddleware(), taskController.AddTask)
	authorized.PUT("/tasks/:id", middlewares.AdminMiddleware(), taskController.UpdateTaskByID)
	authorized.DELETE("/tasks/:id", middlewares.AdminMiddleware(), taskController.DeleteTaskByID)

	// user promotion route
	authorized.POST("/promote", middlewares.AdminMiddleware(), userController.PromoteUser)

	return router
}
