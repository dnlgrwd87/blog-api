package users

import (
	"errors"
	"github.com/dnlgrwd87/blog-api/models"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (service *UserService) GetAllUsers() []models.User {
	var users []models.User
	service.DB.Find(&users)
	return users
}

func (service *UserService) GetUserById(id string) (models.User, error) {
	var user models.User
	if service.DB.Preload("Posts").First(&user, id).RecordNotFound() {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

func (service *UserService) CreateUser(user models.User) (models.User, error) {
	if err := service.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
