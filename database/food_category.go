package database

import (
	"kosmeal/config"
	"kosmeal/model"
)

func GetFoodCategories() []model.FoodCategory {
	var foodcategories []model.FoodCategory
	config.DB.Find(&foodcategories)
	return foodcategories
}

func GetFoodCategoryByID(id string) model.FoodCategory {
	var foodcategory model.FoodCategory
	config.DB.Where("id = ?", id).Find(&foodcategory)
	return foodcategory
}
