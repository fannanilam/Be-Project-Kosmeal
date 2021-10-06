package controller

import (
	"kosmeal/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestGetAllAdminController(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/admin", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)

	err := GetAllAdminController(c)
	if err != nil {
		t.Error(err.Error())
	}

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}
