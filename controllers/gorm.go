package controllers

import "github.com/jinzhu/gorm"

type ConnDB struct {
	DB *gorm.DB
}