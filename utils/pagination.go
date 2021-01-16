package utils

import (
	models "github.com/blackwhite-nabila/gin-app/models"

	"strconv"
	"github.com/gin-gonic/gin"
)

func GeneratePagination(c *gin.Context) models.Pagination {
	limit := 2
	page := 1
	sort := "created_at asc"
	query := c.Request.URL.Query()

	for key, value := range query {
		queryValue := value[len(value)-1]

		switch key {
			case "size":
				limit, _ = strconv.Atoi(queryValue)
				break
			
			case "page":
				page, _ = strconv.Atoi(queryValue)
				break
			case "sort":
				sort = queryValue
				break
		}


	}

	return models.Pagination{
		Limit: limit,
		Page: page,
		Sort: sort,
	}
}