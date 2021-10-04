package model

type Admin struct {
	//gorm.Model
	ID       uint   `gorm:"primarykey"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}
