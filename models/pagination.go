package models

import 	_"github.com/jinzhu/gorm"

type Pagination struct {
	Limit int `json:"limit"`
	Page int `json:"page"`
	Sort string `json:"sort"`
}