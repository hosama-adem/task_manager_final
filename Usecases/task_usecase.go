package Usecases

import (
	"errors"
	"task_manager/Domain"
)

type TaskUsecase interface {
	AddNewTask(task Domain.Task) (Domain.Task, error)
	GetTaskByID(id int) *Domain.Task
	RemoveTaskByID(id int) bool
	UpdateTaskByID(id int, updatedTask Domain.Task) *Domain.Task
	GetAllTasks() []Domain.Task
}

type taskUsecase struct {
	repositories Domain.TaskRepository
}

func NewTaskUsecase(repository Domain.TaskRepository) TaskUsecase {
	return &taskUsecase{
		repositories: repository,
	}
}

func (u *taskUsecase) AddNewTask(task Domain.Task) (Domain.Task, error) {
	if task.Title == "" {
		return Domain.Task{}, errors.New("the title can't not be empty")

	}
	return u.repositories.AddNewTask(task)

}

func (u *taskUsecase) GetTaskByID(id int) *Domain.Task {
	task := u.repositories.GetTaskByID(id)

	if task == nil {
		return nil
	}

	return task
}

func (u *taskUsecase) RemoveTaskByID(id int) bool {
	task := u.repositories.RemoveTaskByID(id)
	return task
}

func (u *taskUsecase) UpdateTaskByID(id int, updatedTask Domain.Task) *Domain.Task {
	return u.repositories.UpdateTaskByID(id, updatedTask)
}

func (u *taskUsecase) GetAllTasks() []Domain.Task {
	return u.repositories.GetAllTasks()
}
