package controller

import (
	"kosmeal/database"
	"kosmeal/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllCommentsController(c echo.Context) error {
	comments := database.GetComments()

	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetAllFoodCommentsController",
		"data":    comments,
	})
}

func CreateCommentController(c echo.Context) error {
	var newComment model.Comment
	if err := c.Bind(&newComment); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateCommentController",
			"error":   err.Error(),
		})
	}

	newComment = database.CreateComment(newComment)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreateCommentController",
		"data":    newComment,
	})
}

func DeleteCommentByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeleteCommentByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "DeleteCommentByIDController",
	})
}

func UpdateCommentByIDController(c echo.Context) error {
	id := c.Param("id")

	var comment model.Comment
	if err := c.Bind(&comment); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "UpdateCommentController",
			"error":   err.Error(),
		})
	}
	database.UpdateCommentByID(id, comment)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "UpdateCommentByIDController",
		"data":    comment,
	})
}
