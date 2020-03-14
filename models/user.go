package models

import "errors"

type User struct {
	//BaseModel
	ID        uint   `json:"id"`
	FirstName string `json:"firstName" gorm:"not null"`
	LastName  string `json:"lastName" gorm:"not null"`
	Email     string `json:"email" gorm:"not null"`
	Posts     []Post `json:"posts" gorm:"foreignkey:UserID"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) ValidateIncomingPayload() error {
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
