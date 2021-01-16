package main

import (
	"github.com/gin-gonic/gin"
	_"github.com/go-sql-driver/mysql"
	"github.com/blackwhite-nabila/gin-app/controllers"
	"github.com/blackwhite-nabila/gin-app/config"
)

func main() {
	// db:= config.Init()

	// connDB := &controllers.ConnDB{DB: db}
	config.Init()

	route := gin.Default()
	route.GET("/furnitures", controllers.GetFurnitures)
	route.GET("/furnitures/:id", controllers.GetFurniture)
	route.POST("/furnitures", controllers.PostFurniture)
	route.PATCH("/furnitures/:id", controllers.UpdateFurniture)
	route.DELETE("/furnitures/:id", controllers.DeleteFurniture)


	route.Run(":3000")


}
