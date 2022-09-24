package models

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	Code        string `gorm:"primarykey"`
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
