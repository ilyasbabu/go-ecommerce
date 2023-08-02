package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
