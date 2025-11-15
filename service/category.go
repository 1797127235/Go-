package service

import (
	"blog-server/global"
	"blog-server/model"

	"gorm.io/gorm"
)

func CreateCategory(category *model.Category) error {
	return global.DB.Create(category).Error
}

// 更新分类
func UpdateCategory(id uint, category *model.Category) error {
	return global.DB.Model(&model.Category{}).Where("id = ?", id).Updates(category).Error
}

func DeleteCategory(id uint) error {
	result := global.DB.Delete(&model.Category{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 { //删除时，如果不存在，返回错误
		return gorm.ErrRecordNotFound
	}
	return nil
}

func GetCategoryByID(id uint) (*model.Category, error) {
	var category model.Category
	if err := global.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// 获取所有分类
func ListCategories() ([]model.Category, error) {
	var categories []model.Category
	if err := global.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
