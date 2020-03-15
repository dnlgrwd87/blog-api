package models

type Post struct {
	ID     int    `json:"id"`
	Title  string `json:"title" gorm:"not null"`
	UserID int    `json:"-"`
}

func (u *Post) TableName() string {
	return "posts"
}
