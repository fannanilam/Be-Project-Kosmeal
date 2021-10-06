package route

import (
	"kosmeal/controller"

	"github.com/labstack/echo/v4"
)

func NewComment(app *echo.Echo) {
	app.GET("/comments", controller.GetAllCommentsController)
	app.POST("/comments", controller.CreateCommentController)
	app.DELETE("/comments/:id", controller.DeleteCommentByIDController)
	app.PUT("/comments/:id", controller.UpdateCommentByIDController)
}
