package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Picture     string
	Name        string
	Description string
	Price       float64
}
