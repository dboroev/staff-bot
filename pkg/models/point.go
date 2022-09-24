package models

import "gorm.io/gorm"

type Point struct {
	gorm.Model
	Name    string
	City    string
	Address string
}
