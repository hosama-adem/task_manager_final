package Usecases

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"

	"task_manager/Domain"
	mocks "task_manager/Mocks"
)

type UserUsecaseSuite struct {
	suite.Suite

	mockRepo *mocks.UserRepositoryMock
	usecase  UserUsecase
}

func (s *UserUsecaseSuite) SetupTest() {
	s.mockRepo = mocks.NewUserRepositoryMock(s.T())
	s.usecase = NewUserUsecase(s.mockRepo)
}

func (s *UserUsecaseSuite) TestCreateUser_Success() {
	user := &Domain.User{
		Email: "hosama@example.com",
	}

	s.mockRepo.EXPECT().
		Create(user).
		Return(user, nil)

	result, err := s.usecase.Create(user)

	s.NoError(err)
	s.Equal(user, result)
}

func (s *UserUsecaseSuite) TestCreateUser_RepositoryError() {
	user := &Domain.User{
		Email: "hosama@example.com",
	}

	repoError := errors.New("repository error")

	s.mockRepo.EXPECT().
		Create(user).
		Return(nil, repoError)

	result, err := s.usecase.Create(user)

	s.Error(err)
	s.Nil(result)
	s.Equal(repoError, err)
}

func (s *UserUsecaseSuite) TestGetByEmail_Success() {
	user := &Domain.User{
		Email: "hosama@example.com",
	}

	s.mockRepo.EXPECT().
		GetByEmail("hosama@example.com").
		Return(user, nil)

	result, err := s.usecase.GetByEmail("hosama@example.com")

	s.NoError(err)
	s.Equal(user, result)
}

func (s *UserUsecaseSuite) TestGetByEmail_NotFound() {
	repoError := errors.New("user not found")

	s.mockRepo.EXPECT().
		GetByEmail("unknown@example.com").
		Return(nil, repoError)

	result, err := s.usecase.GetByEmail("unknown@example.com")

	s.Error(err)
	s.Nil(result)
	s.Equal(repoError, err)
}

func TestUserUsecaseSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseSuite))
}
