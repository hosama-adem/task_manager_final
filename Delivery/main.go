package main

import (
	"task_manager/Delivery/Controllers"
	"task_manager/Delivery/routers"
	"task_manager/Repositories"
	"task_manager/Usecases"
)

// @title Task Manager API
// @version 1.0
// @description REST API for managing tasks.
// @host localhost:8080
// @BasePath /
// @title Task Manager API
// @version 1.0
// @description REST API for managing tasks.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization	
// @Security BearerAuth

func main() {
	taskrepository := Repositories.NewTaskStore()
	userrepository := Repositories.NewUserStore()

	taskusecase := Usecases.NewTaskUsecase(taskrepository)
	userusecase := Usecases.NewUserUsecase(userrepository)

	controller := Controllers.NewTaskController(taskusecase, userusecase)

	r := routers.Router(controller)
	r.Run()

}
