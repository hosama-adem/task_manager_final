package Domain

import (
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type TaskRepository interface {
	AddNewTask(task Task) (Task, error)
	GetAllTasks() []Task
	GetTaskByID(id int) *Task
	UpdateTaskByID(id int, updatedTask Task) *Task
	RemoveTaskByID(id int) bool
}

type UserRepository interface {
	Create(user *User) (*User, error)
	GetByEmail(email string) (*User, error)
}
