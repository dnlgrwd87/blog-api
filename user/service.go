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
	GetUserById(id int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
}

func (service *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if res := service.DB.Find(&users); res.Error != nil {
		return users, res.Error
	}
	return users, nil
}

func (service *UserService) GetUserById(id int) (models.User, error) {
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
