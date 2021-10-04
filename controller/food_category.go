package controller

import (
	"kosmeal/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllFoodCategoriesController(c echo.Context) error {
	foodcategories := database.GetFoodCategories()

	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetAllFoodCategoriesController",
		"data":    foodcategories,
	})
}

func GetFoodCategoryByIDController(c echo.Context) error {
	id := c.Param("id")
	foodcategory := database.GetFoodCategoryByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetFoodCategoryByIDController",
		"data":    foodcategory,
	})
}
