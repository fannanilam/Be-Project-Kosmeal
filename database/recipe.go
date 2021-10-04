package database

import (
	"kosmeal/config"
	"kosmeal/model"
)

func GetRecipes() []model.Recipe {
	var recipes []model.Recipe
	config.DB.Preload("FoodCategory").Find(&recipes)
	return recipes
}

func GetRecipeByID(id string) model.Recipe {
	var recipe model.Recipe
	config.DB.Where("id = ?", id).Find(&recipe)
	return recipe
}

func CreateRecipe(recipe model.Recipe) model.Recipe {
	config.DB.Preload("FoodCategory").Create(&recipe)
	return recipe
}

func DeleteRecipeByID(id string) {
	var recipe model.Recipe
	config.DB.Where("id = ?", id).Delete(&recipe)
}

func UpdateRecipeByID(id string, recipe model.Recipe) {
	config.DB.Where("id = ?", id).Updates(&recipe)
}
