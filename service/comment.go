package service

import (
	"blog-server/global"
	"blog-server/model"
)

func CreateComment(comment *model.Comment) error {
	return global.DB.Create(comment).Error
}


func ListComments(postID uint) ([]model.Comment, error) {
	var comments []model.Comment
	if err := global.DB.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}