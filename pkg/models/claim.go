package models

import (
	"time"

	"gorm.io/gorm"
)

type Claim struct {
	gorm.Model
	Name           string
	JobCode        string
	Job            Job `gorm:"foreignKey:JobCode"`
	PointID        uint
	Point          Point `gorm:"foreignKey:PointID"`
	WorkerAmount   uint
	DatetimeFrom   time.Time
	DatetimeTo     time.Time
	Status         string
	DatetimeFinish time.Time
}
