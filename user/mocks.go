package user

import (
	"errors"
	"github.com/dnlgrwd87/blog-api/models"
)

type MockUserService struct{}

func (service *MockUserService) GetAllUsers() ([]models.User, error) {
	return []models.User{
		{ID: 1, FirstName: "Daniel", LastName: "Garwood", Email: "dnlgrwd@gmail.com"},
	}, nil
}

func (service *MockUserService) GetUserById(id int) (models.UserPostsDTO, error) {
	return models.UserPostsDTO{ID: 1, FirstName: "Daniel", LastName: "Garwood", Email: "dnlgrwd@gmail.com"}, nil
}

func (service *MockUserService) CreateUser(user models.User) (models.User, error) {
	return models.User{ID: 1, FirstName: "Daniel", LastName: "Garwood", Email: "dnlgrwd@gmail.com"}, nil
}

type MockErrorUserService struct{}

func (service *MockErrorUserService) GetAllUsers() ([]models.User, error) {
	return []models.User{}, errors.New("an error occurred")
}

func (service *MockErrorUserService) GetUserById(id int) (models.UserPostsDTO, error) {
	return models.UserPostsDTO{}, errors.New("an error occurred")
}

func (service *MockErrorUserService) CreateUser(user models.User) (models.User, error) {
	return models.User{}, errors.New("an error occurred")
}
