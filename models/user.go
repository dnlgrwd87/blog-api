package models

import "errors"

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName" gorm:"not null"`
	LastName  string `json:"lastName" gorm:"not null"`
	Email     string `json:"email" gorm:"not null"`
	Posts     []Post `json:"posts" gorm:"foreignkey:UserID"`
}

type UserListDTO struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName" gorm:"not null"`
	LastName  string `json:"lastName" gorm:"not null"`
	Email     string `json:"email" gorm:"not null"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) ValidateUserJson() error {
	if u.FirstName == "" {
		return errors.New("body is missing field firstName")
	}
	if u.LastName == "" {
		return errors.New("body is missing field lastName")
	}
	if u.Email == "" {
		return errors.New("body is missing field email")
	}
	return nil
}

func MapToUserListDTO(users []User) []UserListDTO {
	var userList []UserListDTO
	for _, user := range users {
		userList = append(userList, UserListDTO{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		})
	}
	return userList
}
