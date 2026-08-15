package Usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	// "github.com/stretchr/testify/mock"
	"errors"
	"task_manager/Domain"
	mocks "task_manager/Mocks"
)

type TaskUsecaseSuite struct {
	suite.Suite

	mockRepo *mocks.TaskRepositoryMock
	usecase  TaskUsecase
}

func TestAddNewTask(t *testing.T) {
	mockRepo := mocks.NewTaskRepositoryMock(t)
	usecase := NewTaskUsecase(mockRepo)
	task := Domain.Task{
		ID:          1,
		Title:       "Learn Go Testing",
		Description: "Practice Testify and Mockery",
	}
	mockRepo.EXPECT().
		AddNewTask(task).
		Return(task, nil)
	result, err := usecase.AddNewTask(task)
	assert.NoError(t, err)
	assert.Equal(t, task, result)

}

func (s *TaskUsecaseSuite) SetupTest() {
	s.mockRepo = mocks.NewTaskRepositoryMock(s.T())
	s.usecase = NewTaskUsecase(s.mockRepo)
}

func (s *TaskUsecaseSuite) TestAddNewTask_EmptyTitle() {
	task := Domain.Task{
		ID:          2,
		Title:       "",
		Description: "This should fail",
	}

	result, err := s.usecase.AddNewTask(task)

	s.Error(err)
	s.Equal(Domain.Task{}, result)
}

func (s *TaskUsecaseSuite) TestAddNewTask_RepositoryError() {
	task := Domain.Task{
		ID:          3,
		Title:       "Learn Mockery",
		Description: "Testing repository errors",
	}

	repositoryError := errors.New("repository error")

	s.mockRepo.EXPECT().
		AddNewTask(task).
		Return(Domain.Task{}, repositoryError)

	result, err := s.usecase.AddNewTask(task)

	s.Error(err)
	s.Equal(repositoryError, err)
	s.Equal(Domain.Task{}, result)
}

// to get task by id
func (s *TaskUsecaseSuite) TaskRepositoryMock_GetTaskByID_Call() {
	task := &Domain.Task{
		ID:          1,
		Title:       "Learn Go",
		Description: "Learn unit testing",
	}

	s.mockRepo.EXPECT().
		GetTaskByID(1).
		Return(task)

	result := s.usecase.GetTaskByID(1)

	s.NotNil(result)
	s.Equal(task, result)
}

// To test get task by id not found check
func (s *TaskUsecaseSuite) TestGetTaskByID_NotFound() {
	s.mockRepo.EXPECT().
		GetTaskByID(999).
		Return(nil)

	result := s.usecase.GetTaskByID(999)

	s.Nil(result)
}

// to remove task by id
func (s *TaskUsecaseSuite) TestRemoveTaskByID_Success() {
	s.mockRepo.EXPECT().
		RemoveTaskByID(1).
		Return(true)

	result := s.usecase.RemoveTaskByID(1)

	s.True(result)
}

func (s *TaskUsecaseSuite) TestRemoveTaskByID_NotFound() {
	s.mockRepo.EXPECT().
		RemoveTaskByID(999).
		Return(false)

	result := s.usecase.RemoveTaskByID(999)

	s.False(result)
}

func (s *TaskUsecaseSuite) TestUpdateTaskByID_Success() {
	updatedTask := Domain.Task{
		ID:          1,
		Title:       "Updated Task",
		Description: "Updated description",
	}

	s.mockRepo.EXPECT().
		UpdateTaskByID(1, updatedTask).
		Return(&updatedTask)

	result := s.usecase.UpdateTaskByID(1, updatedTask)

	s.NotNil(result)
	s.Equal(&updatedTask, result)
}

func (s *TaskUsecaseSuite) TestUpdateTaskByID_NotFound() {
	updatedTask := Domain.Task{
		ID:    999,
		Title: "Updated Task",
	}
	s.mockRepo.EXPECT().
		UpdateTaskByID(999, updatedTask).
		Return(nil)

	result := s.usecase.UpdateTaskByID(999, updatedTask)

	s.Nil(result)
}

func (s *TaskUsecaseSuite) TestGetAllTasks() {
	tasks := []Domain.Task{
		{
			ID:          1,
			Title:       "Task One",
			Description: "First task",
		},
		{
			ID:          2,
			Title:       "Task Two",
			Description: "Second task",
		},
	}

	s.mockRepo.EXPECT().
		GetAllTasks().
		Return(tasks)

	result := s.usecase.GetAllTasks()

	s.Equal(tasks, result)
}

func TestTaskUsecaseSuite(t *testing.T) {
	suite.Run(t, new(TaskUsecaseSuite))
}
