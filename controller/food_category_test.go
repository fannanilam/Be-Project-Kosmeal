package controller

import (
	"kosmeal/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestGetAllFoodCategoriesController(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/foodcategories", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)

	err := GetAllFoodCategoriesController(c)
	if err != nil {
		t.Error(err.Error())
	}

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}

func TestGetFoodCategoryByIDController(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "/foodcategories/:id", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)

	err := GetFoodCategoryByIDController(c)
	if err != nil {
		t.Error(err.Error())
	}

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}
