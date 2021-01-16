package models

import (
	"github.com/jinzhu/gorm"

)

type (
	Furniture struct {
		gorm.Model
		Name string `json:"name" binding:"required"`
		Description string `json:"description" binding:"required"`
		Harga int `json:"harga" binding:"required"`
	}

	TransformedFurniture struct {
		ID uint `json:"id"`
		Name string `json:"name"`
		Description string `json:"description"`
		Harga int `json:"harga"`

	}

)