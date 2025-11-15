package service

import (
	"blog-server/global"
	"blog-server/model"

	"gorm.io/gorm"
)

// 创建文章
func CreatePost(p *model.Post) error {
	return global.DB.Create(p).Error
}

// 根据文章ID获取文章
func GetPostByID(id uint) (*model.Post, error) {
	var post model.Post
	if err := global.DB.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

// 分页获取文章列表
func ListPosts(page, pageSize int) ([]model.Post, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	var posts []model.Post
	var total int64

	db := global.DB.Model(&model.Post{})

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := db.Order("created_at DESC").Limit(pageSize).Offset(offset).Find(&posts).Error; err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

// 更新文章
func UpdatePost(id uint, p *model.Post) error {
	return global.DB.Model(&model.Post{}).Where("id = ?", id).Updates(p).Error
}

// 删除文章
func DeletePost(id uint) error {
	result := global.DB.Delete(&model.Post{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
