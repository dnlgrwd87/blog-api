package post

import (
	"github.com/jinzhu/gorm"
)

type PostService struct {
	DB *gorm.DB
}
