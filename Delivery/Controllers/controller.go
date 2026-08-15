package Controllers

import (
	"errors"
	"net/http"
	"strconv"
	"task_manager/Domain"
	"task_manager/Usecases"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskUsecase Usecases.TaskUsecase
	userUsecase Usecases.UserUsecase
}

func NewTaskController(
	taskUsecase Usecases.TaskUsecase,
	userUsecase Usecases.UserUsecase,
) *TaskController {
	return &TaskController{
		taskUsecase: taskUsecase,
		userUsecase: userUsecase,
	}
}

// to get all tasks
func (d *TaskController) GetAllTasks(c *gin.Context) {
	tasks := d.taskUsecase.GetAllTasks()
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// to get tasks by id
// To Get Task by ID
func (d *TaskController) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := d.taskUsecase.GetTaskByID(idInt)
	if task == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"task": task})
	}

}

// To Update Task by ID
func (d *TaskController) UpdateTaskByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedTask Domain.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	task := d.taskUsecase.UpdateTaskByID(idInt, updatedTask)
	if task == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("Task not found")})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task": task})
}

// To Remove Task by ID
func (d *TaskController) RemoveTaskByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := d.taskUsecase.RemoveTaskByID(idInt)
	if !data {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("Task not found")})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task removed successfully"})
}

// To add new Task
func (d *TaskController) AddTask(c *gin.Context) {
	var newTask Domain.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	newtask, err := d.taskUsecase.AddNewTask(newTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"task": newtask})
}

func (d *TaskController) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := d.userUsecase.Login(
		req.Email,
		req.Password,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token":      token,
		"type":       "Barrer",
		"expires_in": 86400,
	})

}

func (d *TaskController) Register(c *gin.Context) {
	var user Domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := d.userUsecase.Register(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// createdUser.Password = ""
	c.JSON(http.StatusCreated, gin.H{"user": createdUser})
}
