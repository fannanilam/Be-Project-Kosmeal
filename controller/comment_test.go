package controller

import (
	"kosmeal/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestGetAllCommentsController(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/comments", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)

	err := GetAllCommentsController(c)
	if err != nil {
		t.Error(err.Error())
	}

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}

func TestCreateCommentController(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/comments", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)

	err := CreateCommentController(c)
	if err != nil {
		t.Error(err.Error())
	}

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}

func TestUpdateCommentByIDController(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "/comments:id", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)

	err := UpdateCommentByIDController(c)
	if err != nil {
		t.Error(err.Error())
	}

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}

func TestDeleteCommentByIDController(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/comments/:id", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)

	err := DeleteCommentByIDController(c)
	if err != nil {
		t.Error(err.Error())
	}

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}
