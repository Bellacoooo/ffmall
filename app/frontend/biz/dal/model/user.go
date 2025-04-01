package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    *string // A pointer to a string, allowing for null values
	Password string  // A regular string field

}
