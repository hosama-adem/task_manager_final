package Usecases

import (
	"errors"
	"task_manager/Domain"
	"task_manager/Infrastructure"
)

type UserUsecase interface {
	Create(user *Domain.User) (*Domain.User, error)
	GetByEmail(email string) (*Domain.User, error)
	Login(email, password string) (string, error)
	Register(user *Domain.User) (*Domain.User, error)
}

type userUsecase struct {
	repositories Domain.UserRepository
}

func NewUserUsecase(repository Domain.UserRepository) *userUsecase {
	return &userUsecase{
		repositories: repository,
	}
}

func (u *userUsecase) Create(user *Domain.User) (*Domain.User, error) {
	return u.repositories.Create(user)
}

func (u *userUsecase) GetByEmail(email string) (*Domain.User, error) {
	return u.repositories.GetByEmail(email)
}

func (u *userUsecase) Login(email, password string) (string, error) {
	user, err := u.repositories.GetByEmail(email)
	if err != nil {
		return "", err
	}

	if ok := Infrastructure.ComparePassword(user.Password, password); !ok {
		return "", errors.New("invalid password")
	}

	token, err := Infrastructure.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *userUsecase) Register(user *Domain.User) (*Domain.User, error) {
	hashedPassword, err := Infrastructure.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword

	if user.Role == "" {
		user.Role = "user"
	}

	return u.repositories.Create(user)
}
