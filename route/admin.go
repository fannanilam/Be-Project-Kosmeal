package route

import (
	"kosmeal/controller"

	"github.com/labstack/echo/v4"
)

func NewAdmin(app *echo.Echo) {
	app.GET("/admin", controller.GetAllAdminController)
}
