package services

import (
	"github.com/invazions/MSSQL_INTERACTION/models"
	"github.com/invazions/MSSQL_INTERACTION/repositories"
)

type UserService interface {
	CreateUser(user models.User) error
	GetUserByID(id int) (models.User, error)
	DeleteUser(id int) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user models.User) error {
	return s.repo.Create(user)
}

func (s *userService) GetUserByID(id int) (models.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}
