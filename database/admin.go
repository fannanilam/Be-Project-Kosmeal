package database

import (
	"kosmeal/config"
	"kosmeal/model"
)

func GetAllAdmin() []model.Admin {
	var admin []model.Admin
	config.DB.Find(&admin)
	return admin
}
