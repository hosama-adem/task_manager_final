package routers

import (
	"task_manager/Delivery/Controllers"
	"task_manager/Infrastructure"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "task_manager/docs"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
}

func Router(controller *Controllers.TaskController) *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)

	router.GET("/tasks", Infrastructure.AuthMiddleware(), controller.GetAllTasks)
	router.GET("/tasks/:id", Infrastructure.AuthMiddleware(), controller.GetTaskByID)
	router.PUT("/tasks/:id", Infrastructure.AuthMiddleware(), controller.UpdateTaskByID)
	router.POST("/tasks", Infrastructure.AuthMiddleware(), controller.AddTask)

	router.DELETE("/tasks/:id", Infrastructure.AuthMiddleware(), Infrastructure.RequireRole("admin"), controller.RemoveTaskByID)

	return router
}
