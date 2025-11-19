package service

import (
	"blog-server/global"
	"blog-server/model"

	"gorm.io/gorm"
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

// 根据评论ID获取评论
func GetCommentByID(id uint) (*model.Comment, error) {
	var comment model.Comment
	if err := global.DB.First(&comment, id).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

// 根据ID删除评论
func DeleteComment(id uint) error {
	result := global.DB.Delete(&model.Comment{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
