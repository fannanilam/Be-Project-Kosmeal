package route

import (
	"kosmeal/controller"

	"github.com/labstack/echo/v4"
)

func NewAdmin(app *echo.Echo) {
	app.GET("/admins", controller.GetAllAdminController)

}
