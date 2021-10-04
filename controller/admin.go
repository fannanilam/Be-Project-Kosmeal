package controller

import (
	"kosmeal/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllAdminController(c echo.Context) error {
	admin := database.GetAllAdmin()

	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetAllAdminController",
		"data":    admin,
	})
}
