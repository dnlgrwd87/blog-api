package models

import (
	"time"
)

type BaseModel struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
}
