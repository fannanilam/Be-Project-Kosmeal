package model

import "time"

type Recipe struct {
	ID             uint          `gorm:"primarykey"`
	FoodCategory   *FoodCategory `json:"foodcategory,omitempty"`
	FoodCategoryID uint          `json:"foodcategoryid"`
	Title          string        `json:"title"`
	Infocalori     string        `json:"infocalori"`
	Ingredients    string        `json:"ingredients"`
	Tools          string        `json:"tools"`
	Method         string        `json:"method"`
	NutritionURL   string        `json:"nutrition_url"`
	Nutritions     []*Nutrition  `json:"nutritions" gorm:"-"`
	CreatedAt      time.Time     `json:"createdat"`
	Comments       []*Comment    `json:"comments"`
}
