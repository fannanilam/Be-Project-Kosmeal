package model

type FoodCategory struct {
	ID           uint      `gorm:"primarykey"`
	FoodCategory string    `json:"foodcategory"`
	Recipes      []*Recipe `json:"recipes,omitempty"`
}
