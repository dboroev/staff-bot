package models

import (
	"gorm.io/gorm"
)

type UserClaim struct {
	gorm.Model
	UserID  uint
	User    User `gorm:"foreignKey:UserID"`
	ClaimID uint
	Claim   Claim `gorm:"foreignKey:ClaimID"`
	Grade   uint
}
