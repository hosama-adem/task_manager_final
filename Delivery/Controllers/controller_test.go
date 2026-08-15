package Controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	// "task_manager/Delivery/Controllers"
	"bytes"
	"encoding/json"
	"task_manager/Domain"
	mocks "task_manager/Mocks"
)

type ControllerSuite struct {
	suite.Suite

	mockUsecase     *mocks.TaskUsecaseMock
	mockUserUsecase *mocks.UserUsecaseMock
	controller      *TaskController
	router          *gin.Engine
}

func (s *ControllerSuite) SetupTest() {
	gin.SetMode(gin.TestMode)

	s.mockUsecase = mocks.NewTaskUsecaseMock(s.T())
	s.mockUserUsecase = mocks.NewUserUsecaseMock(s.T())
	s.controller = NewTaskController(s.mockUsecase, s.mockUserUsecase)

	s.router = gin.New()

	// Register your controller routes here according to your controller.
	s.router.GET("/tasks", s.controller.GetAllTasks)
	s.router.GET("/tasks/:id", s.controller.GetTaskByID)
	s.router.POST("/tasks", s.controller.AddTask)
	s.router.PUT("/tasks/:id", s.controller.UpdateTaskByID)
	s.router.DELETE("/tasks/:id", s.controller.RemoveTaskByID)
}

func (s *ControllerSuite) TestGetAllTasks_Success() {
	tasks := []Domain.Task{
		{
			ID:          1,
			Title:       "Learn Go",
			Description: "Learn testing",
		},
		{
			ID:          2,
			Title:       "Practice Mockery",
			Description: "Write mocks",
		},
	}

	s.mockUsecase.EXPECT().
		GetAllTasks().
		Return(tasks)

	req := httptest.NewRequest(
		http.MethodGet,
		"/tasks",
		nil,
	)

	rec := httptest.NewRecorder()

	s.router.ServeHTTP(rec, req)

	s.Equal(http.StatusOK, rec.Code)
}

func (s *ControllerSuite) TestGetTaskByID_Success() {
	task := &Domain.Task{
		ID:          1,
		Title:       "Learn Go",
		Description: "Learn testing",
	}

	s.mockUsecase.EXPECT().
		GetTaskByID(1).
		Return(task)

	req := httptest.NewRequest(
		http.MethodGet,
		"/tasks/1",
		nil,
	)

	rec := httptest.NewRecorder()

	s.router.ServeHTTP(rec, req)

	s.Equal(http.StatusOK, rec.Code)
}

func (s *ControllerSuite) TestGetTaskByID_NotFound() {
	s.mockUsecase.EXPECT().
		GetTaskByID(999).
		Return(nil)

	req := httptest.NewRequest(
		http.MethodGet,
		"/tasks/999",
		nil,
	)

	rec := httptest.NewRecorder()

	s.router.ServeHTTP(rec, req)

	s.Equal(http.StatusNotFound, rec.Code)
}

func (s *ControllerSuite) TestAddTask_Success() {
	task := Domain.Task{
		ID:          1,
		Title:       "Learn Go",
		Description: "Learn controller testing",
	}

	s.mockUsecase.EXPECT().
		AddNewTask(task).
		Return(task, nil)

	body, err := json.Marshal(task)
	s.NoError(err)

	req := httptest.NewRequest(
		http.MethodPost,
		"/tasks",
		bytes.NewBuffer(body),
	)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	s.router.ServeHTTP(rec, req)

	s.Equal(http.StatusCreated, rec.Code)
	s.Contains(rec.Body.String(), "Learn Go")
}

func (s *ControllerSuite) TestAddTask_InvalidJSON() {
	body := []byte(`{"title": "Learn Go"`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/tasks",
		bytes.NewBuffer(body),
	)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	s.router.ServeHTTP(rec, req)

	s.Equal(http.StatusBadRequest, rec.Code)
}

func (s *ControllerSuite) TestUpdateTask_Success() {
	task := Domain.Task{
		ID:          1,
		Title:       "Updated Go",
		Description: "Updated description",
	}

	s.mockUsecase.EXPECT().
		UpdateTaskByID(1, task).
		Return(&task)

	body, err := json.Marshal(task)
	s.NoError(err)

	req := httptest.NewRequest(
		http.MethodPut,
		"/tasks/1",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	s.router.ServeHTTP(rec, req)

	s.Equal(http.StatusOK, rec.Code)
}

func (s *ControllerSuite) TestUpdateTask_InvalidID() {
	task := &Domain.Task{
		Title:       "Updated Go",
		Description: "Updated description",
	}

	body, err := json.Marshal(task)
	s.NoError(err)

	req := httptest.NewRequest(
		http.MethodPut,
		"/tasks/abc",
		bytes.NewBuffer(body),
	)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	s.router.ServeHTTP(rec, req)

	s.Equal(http.StatusBadRequest, rec.Code)
}

func (s *ControllerSuite) TestDeleteTask_InvalidID() {
	req := httptest.NewRequest(
		http.MethodDelete,
		"/tasks/abc",
		nil,
	)

	rec := httptest.NewRecorder()

	s.router.ServeHTTP(rec, req)

	s.Equal(http.StatusBadRequest, rec.Code)
}

func TestControllerSuite(t *testing.T) {
	suite.Run(t, new(ControllerSuite))
}
