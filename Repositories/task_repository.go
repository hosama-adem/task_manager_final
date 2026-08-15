package Repositories

import (
	"errors"
	"task_manager/Domain"
)

type TaskStore struct {
	tasks []Domain.Task
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks: make([]Domain.Task, 0),
	}
}

// To Add New Task
func (s *TaskStore) AddNewTask(task Domain.Task) (Domain.Task, error) {
	taskid := task.ID

	for _, t := range s.tasks {
		if t.ID == taskid {
			return Domain.Task{}, errors.New("Task with this ID already exists")
		}
	}

	s.tasks = append(s.tasks, task)
	return task, nil
}

// getting Task by Id
func (s *TaskStore) GetTaskByID(id int) *Domain.Task {
	for i, _ := range s.tasks {
		if s.tasks[i].ID == id {
			return &s.tasks[i]
		}
	}
	return nil
}

// Remove the Task by id
func (s *TaskStore) RemoveTaskByID(id int) bool {
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return true
		}
	}
	return false

}

// update the taskby id
func (s *TaskStore) UpdateTaskByID(id int, updatedTask Domain.Task) *Domain.Task {
	for i, task := range s.tasks {
		if task.ID == id {
			if updatedTask.Title != "" {
				s.tasks[i].Title = updatedTask.Title
			}

			if updatedTask.Description != "" {
				s.tasks[i].Description = updatedTask.Description
			}
			return &s.tasks[i]

		}

	}
	return nil
}

// to get All tasks
func (s *TaskStore) GetAllTasks() []Domain.Task {
	return s.tasks
}
