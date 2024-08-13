package router

import (
	"task_managment_api/delivery/controllers"
	"task_managment_api/infrastructure"
	"task_managment_api/repositories"
	"task_managment_api/usecases"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(db *mongo.Database, timeOut time.Duration) *gin.Engine {

	tr := repositories.NewTaskRepository(db)
	tc := repositories.NewUserRepository(db)
	
	taskController := controllers.NewTaskController(usecases.NewTaskUsecase(tr, timeOut)) 
	userController := controllers.NewUserController(usecases.NewUserUsecase(tc, timeOut))

	router := gin.Default()

	// public routes
	router.POST("/register", userController.RegisterUser)
	router.POST("/login", userController.LoginUser)

	// private routes
	authorized := router.Group("/")
	authorized.Use(infrastructure.AuthMiddleware())

	// task routes
	authorized.GET("/tasks", taskController.GetTasks)
	authorized.GET("/tasks/:id", taskController.GetTaskByID)
	authorized.POST("/tasks", infrastructure.AdminMiddleware(), taskController.CreateTask)
	authorized.PUT("/tasks/:id", infrastructure.AdminMiddleware(), taskController.UpdateTaskByID)
	authorized.DELETE("/tasks/:id", infrastructure.AdminMiddleware(), taskController.DeleteTaskByID)

	// user promotion route
	authorized.POST("/promote", infrastructure.AdminMiddleware(), userController.PromoteUser)

	return router
}
