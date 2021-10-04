package model

import "time"

type Comment struct {
	ID        uint      `gorm:"primarykey"`
	RecipeID  uint      `json:"recipeid"`
	UserID    uint      `json:"userid"`
	User      *User     `json:"user,omitempty"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"createdat"`
	Recipe    *Recipe   `json:"recipes,omitempty"`
}
