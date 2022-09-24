package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Firstname  string
	Secondname string
	Lastname   string
	Phone      string
	Gender     string
}
