package models

import (
	"errors"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName" gorm:"not null"`
	LastName  string `json:"lastName" gorm:"not null"`
	Email     string `json:"email" gorm:"not null"`
}

type UserPostsDTO struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName" gorm:"not null"`
	LastName  string `json:"lastName" gorm:"not null"`
	Email     string `json:"email" gorm:"not null"`
	Posts     []Post `json:"posts" gorm:"foreignkey:UserID"`
}

func (user *UserPostsDTO) TableName() string {
	return "users"
}

func (user *User) ValidateUserJson() error {
	if user.FirstName == "" {
		return errors.New("body is missing field firstName")
	}
	if user.LastName == "" {
		return errors.New("body is missing field lastName")
	}
	if user.Email == "" {
		return errors.New("body is missing field email")
	}
	return nil
}
