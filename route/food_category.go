package route

import (
	"kosmeal/controller"

	"github.com/labstack/echo/v4"
)

func NewFoodCategory(app *echo.Echo) {
	app.GET("/foodcategories", controller.GetAllFoodCategoriesController)
	app.GET("/foodcategories/:id", controller.GetFoodCategoryByIDController)
}
