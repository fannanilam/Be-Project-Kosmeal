package database

import (
	"kosmeal/config"
	"kosmeal/model"
)

func GetComments() []model.Comment {
	var comments []model.Comment
	config.DB.Find(&comments)
	return comments
}

func CreateComment(comment model.Comment) model.Comment {
	config.DB.Create(&comment)
	return comment
}

func DeleteCommentByID(id string) {
	var comment model.Comment
	config.DB.Where("id = ?", id).Delete(&comment)
}

func UpdateCommentByID(id string, comment model.Comment) {
	config.DB.Where("id = ?", id).Updates(&comment)
}
