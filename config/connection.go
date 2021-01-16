package config


import (
	
	"github.com/jinzhu/gorm"
	"github.com/blackwhite-nabila/gin-app/models"
)


var Db *gorm.DB
func Init() *gorm.DB {

	var err error
	Db,err = gorm.Open("mysql", "root:root123@tcp(localhost:3306)/furniture?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database !")
	}

	Db.AutoMigrate(models.Furniture{})

	return Db

}