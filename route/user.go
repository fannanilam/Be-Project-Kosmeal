package route

import (
	"kosmeal/controller"
	"kosmeal/middleware"

	"github.com/labstack/echo/v4"
)

func NewUser(app *echo.Echo) {
	app.GET("/users", controller.GetAllUsersController, middleware.JWTAuthMiddleware)
	app.POST("/users", controller.CreateUserController)
	app.GET("/users/:id", controller.GetUserByIDController)
	app.DELETE("/users/:id", controller.DeleteUserByIDController)
	app.PUT("/users/:id", controller.UpdateUserByIDController)
}
