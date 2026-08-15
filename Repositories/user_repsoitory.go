package Repositories

import (
	"errors"
	// "net/http"
	"sync"
	"task_manager/Domain"
	// "github.com/gin-gonic/gin"
	// "golang.org/x/crypto/bcrypt"
)

type UserStore struct {
	users  map[string]*Domain.User
	mu     sync.RWMutex
	nextID uint
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func NewUserStore() *UserStore {
	return &UserStore{
		users:  make(map[string]*Domain.User),
		nextID: 1,
	}
}

func (s *UserStore) Create(user *Domain.User) (*Domain.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[user.Email]; exists {
		return nil, errors.New("user already exists")
	}

	user.ID = int(s.nextID)
	s.nextID++
	s.users[user.Email] = user
	return user, nil
}

func (s *UserStore) GetByEmail(email string) (*Domain.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	user, exists := s.users[email]
	if exists {
		return user, nil
	}
	return nil, errors.New("user not found")
}
