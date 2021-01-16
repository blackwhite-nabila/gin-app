package repositories

import (
	models "github.com/blackwhite-nabila/gin-app/models"
	_"fmt"
	_"strconv"
	config "github.com/blackwhite-nabila/gin-app/config"

)


func GetAll(furniture *models.Furniture, pagination *models.Pagination) (*[]models.Furniture, error) {
	
	var furnitures []models.Furniture
	

	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := config.Db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuilder.Model(&models.Furniture{}).Where(furniture).Find(&furnitures)


	if result.Error != nil {
		message := result.Error
		return nil, message
	}
	
	return &furnitures, nil
}

func GetById(id int) (*[]models.Furniture, error) {

	var furnitures []models.Furniture
	


	result := config.Db.Model(&models.Furniture{}).First(&furnitures, id)

	if result.Error != nil {
		message := result.Error
		return nil, message
	}

	
	return &furnitures, nil
}

func Add(furniture *models.Furniture) (*[]models.Furniture, error) {
	
	config.Db.Create(&furniture)
	
	result, _ := GetById(int(furniture.ID))

	return result, nil
}

func Update(id int, furniture *models.Furniture) (*[]models.Furniture, error) {
	
	queryBuilder := config.Db.Model(&models.Furniture{}).Where("id = ?", id)

	result := queryBuilder.Updates(furniture)

	if result.Error != nil {
		message := result.Error
		return nil, message
	}	

	furnitures, _ := GetById(id) 	

	
	return furnitures, nil
}

func Remove(id int) (*[]models.Furniture, error) {

	var furnitures []models.Furniture


	queryBuilder := config.Db.Model(&models.Furniture{}).Where("id = ?", id)
	result := queryBuilder.Delete(&furnitures)

	if result.Error != nil {
		message := result.Error
		return nil, message
	}	

	config.Db.Unscoped().Where("id = ?", id).Find(&furnitures)

	
	return &furnitures, nil

}