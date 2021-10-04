package controller

import (
	"kosmeal/database"
	"kosmeal/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllRecipesController(c echo.Context) error {
	recipes := database.GetRecipes()

	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetAllRecipesController",
		"data":    recipes,
	})
}

func GetRecipeByIDController(c echo.Context) error {
	id := c.Param("id")
	recipe := database.GetRecipeByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetRecipeByIDController",
		"data":    recipe,
	})
}

func DeleteRecipeByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeleteRecipeByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "DeleteRecipeByIDController",
	})
}

func UpdateRecipeByIDController(c echo.Context) error {
	id := c.Param("id")

	var recipe model.Recipe
	if err := c.Bind(&recipe); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateRecipeController",
			"error":   err.Error(),
		})
	}
	database.UpdateRecipeByID(id, recipe)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetRecipeByIDController",
		"data":    recipe,
	})
}

func CreateRecipeController(c echo.Context) error {
	var newRecipe model.Recipe
	if err := c.Bind(&newRecipe); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateRecipeController",
			"error":   err.Error(),
		})
	}

	newRecipe = database.CreateRecipe(newRecipe)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreateRecipeController",
		"data":    newRecipe,
	})
}
