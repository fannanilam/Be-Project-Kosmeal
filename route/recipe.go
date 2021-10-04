package route

import (
	"kosmeal/controller"

	"github.com/labstack/echo/v4"
)

func NewRecipe(app *echo.Echo) {
	app.GET("/recipes", controller.GetAllRecipesController)
	app.GET("/recipes/:id", controller.GetRecipeByIDController)
	app.POST("/recipes", controller.CreateRecipeController)
	app.DELETE("/recipes/:id", controller.DeleteRecipeByIDController)
	app.PUT("/recipes/:id", controller.UpdateRecipeByIDController)
}
