package models

type Post struct {
	ID     uint   `json:"id"`
	Title  string `json:"title" gorm:"not null"`
	UserID uint   `json:"-"`
}

func (u *Post) TableName() string {
	return "posts"
}
