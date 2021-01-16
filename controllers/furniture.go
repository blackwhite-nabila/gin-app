package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/blackwhite-nabila/gin-app/models"
	_"fmt"
	"strconv"
	repo "github.com/blackwhite-nabila/gin-app/repositories"
	utils "github.com/blackwhite-nabila/gin-app/utils"


)

func  GetFurnitures(c *gin.Context) {
	var (
		raw models.Furniture
		furnitures []models.TransformedFurniture
		result gin.H
	)


	pagination := utils.GeneratePagination(c)
	furnituresList, err := repo.GetAll(&raw, &pagination)

	if err != nil {
		result = gin.H{
			"status" : 500,
			"data" : nil,
			"message" : err,
		}

	}

	for _, data := range *furnituresList {
				furnitures = append(furnitures, models.TransformedFurniture{
					ID: data.ID,
					Name: data.Name,
					Description: data.Description,
					Harga: data.Harga,
				})
			}

	result = gin.H{
		"status": 200,
		"data" : furnitures,
		"page" : pagination.Page,
		"size" : pagination.Limit,
		"count" : len(furnitures),
	}

	c.JSON(http.StatusOK, result)
}

func GetFurniture(c *gin.Context)  {
	
	var (
		result gin.H
		furniture []models.TransformedFurniture
	)
	fID, _ := strconv.Atoi(c.Param("id"))

	furnitureList, err := repo.GetById(fID)

	if err != nil {
		result = gin.H{
			"status" : 500,
			"data" : nil,
			"message" : err,
		}
	}

	for _, data := range *furnitureList {
		furniture = append(furniture, models.TransformedFurniture{
			ID: data.ID,
			Name: data.Name,
			Description: data.Description,
			Harga: data.Harga,
		})		
	}

	result = gin.H{
		"status" : 200,
		"data" : furniture,
		"count" : len(furniture),
	}

	c.JSON(http.StatusOK, result)

}

func PostFurniture(c *gin.Context)  {
	var (
		result gin.H
		req *models.Furniture
		furniture []models.TransformedFurniture

	)
	c.Header("Content-Type","application/json")

	c.Bind(&req)


	form := &models.Furniture{
		Name: req.Name,
		Description: req.Description,
		Harga: req.Harga,
	}

	furnitureList, err := repo.Add(form)

	if err != nil {
		result = gin.H{
			"status" : 500,
			"data" : nil,
			"message" : err,
		}
		return
	}

	for _, data := range *furnitureList {
		furniture = append(furniture, models.TransformedFurniture{
			ID: data.ID,
			Name: data.Name,
			Description: data.Description,
			Harga: data.Harga,
		})		
	}

	result = gin.H{
		"status": 200,
		"data": furniture,
		"count": len(furniture),
	}

	c.JSON(http.StatusOK, result)

}

func UpdateFurniture(c *gin.Context)  {
	var (
		result gin.H
		req *models.Furniture
		furniture []models.TransformedFurniture

	)
	c.Header("Content-Type","application/json")

	c.Bind(&req)
	fID, _ := strconv.Atoi(c.Param("id"))

	
	form := &models.Furniture{
		Name: req.Name,
		Description: req.Description,
		Harga: req.Harga,
	}

	furnitureList, err := repo.Update(fID, form)

	if err != nil {
		result = gin.H{
			"status" : 500,
			"data" : nil,
			"message" : err,
		}
		return
	}

	for _, data := range *furnitureList {
		furniture = append(furniture, models.TransformedFurniture{
			ID: data.ID,
			Name: data.Name,
			Description: data.Description,
			Harga: data.Harga,
		})		
	}

	result = gin.H{
		"status": 200,
		"data": furniture,
		"count": len(furniture),
	}

	c.JSON(http.StatusOK, result)

}

func DeleteFurniture(c *gin.Context)  {
	var (
		result gin.H
		furniture []models.TransformedFurniture
	)
	fID, _ := strconv.Atoi(c.Param("id"))

	furnitureList, err := repo.Remove(fID)

	if err != nil {
		result = gin.H{
			"status" : 500,
			"data" : nil,
			"message" : err,
		}
	}

	for _, data := range *furnitureList {
		furniture = append(furniture, models.TransformedFurniture{
			ID: data.ID,
			Name: data.Name,
			Description: data.Description,
			Harga: data.Harga,
		})		
	}

	result = gin.H{
		"status" : 200,
		"data" : furniture,
		"count" : len(furniture),
	}

	c.JSON(http.StatusOK, result)
}