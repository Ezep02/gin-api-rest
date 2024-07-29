package services

import (
	"github.com/go-api-rest/internal/models"
	"github.com/go-api-rest/internal/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(UserRepo *repositories.UserRepository) *UserService {
	return &UserService{
		repo: UserRepo,
	}
}

func (us *UserService) GetUser(id int) (*models.User, error) {
	return us.repo.FindByID(id)
}

func (us *UserService) AddNewUser(user *models.User) (*models.User, error) {
	return us.repo.AddNewUser(user)
}
