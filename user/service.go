package user

import (
	"errors"
	"github.com/dnlgrwd87/blog-api/models"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	DB *gorm.DB
}

type UserServiceInterface interface {
	GetAllUsers() ([]models.User, error)
	GetUserById(id int) (models.UserPostsDTO, error)
	CreateUser(user models.User) (models.User, error)
}

func (service *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if res := service.DB.Find(&users); res.Error != nil {
		return []models.User{}, res.Error
	}

	return users, nil
}

func (service *UserService) GetUserById(id int) (models.UserPostsDTO, error) {
	var user models.UserPostsDTO
	if service.DB.Preload("Posts").First(&user, id).RecordNotFound() {
		return models.UserPostsDTO{}, errors.New("user not found")
	}
	return user, nil
}

func (service *UserService) CreateUser(user models.User) (models.User, error) {
	if err := service.DB.Create(&user).Error; err != nil {
		return models.User{}, errors.New("error creating user")
	}

	return user, nil
}
