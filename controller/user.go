package controller

import (
	"kosmeal/database"
	"kosmeal/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllUsersController(c echo.Context) error {
	users := database.GetUsers()

	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetAllUsersController",
		"data":    users,
	})
}

func GetUserByIDController(c echo.Context) error {
	id := c.Param("id")
	user := database.GetUserByID(id)
	user.Password = ""
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetUserByIDController",
		"data":    user,
	})
}

func DeleteUserByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeleteUserByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "DeleteUserByIDController",
	})
}

func UpdateUserByIDController(c echo.Context) error {
	id := c.Param("id")

	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateUserController",
			"error":   err.Error(),
		})
	}
	database.UpdateUserByID(id, user)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetUserByIDController",
		"data":    user,
	})
}

func CreateUserController(c echo.Context) error {
	var newUser model.User
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateUserController",
			"error":   err.Error(),
		})
	}

	newUser = database.CreateUser(newUser)
	newUser.Password = ""
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreateUserController",
		"data":    newUser,
	})
}
