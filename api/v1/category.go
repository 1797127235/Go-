package v1

import (
	"blog-server/model"
	"blog-server/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	Description string `json:"description"`
}

// 创建分类
func CreateCategory(c *gin.Context) {
	// 从上下文获取当前用户，检查是否为管理员
	uidVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not logged in"})
		return
	}
	uid, ok := uidVal.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user id in context"})
		return
	}

	user, err := service.GetUserByID(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user.Role != 1 {
		c.JSON(http.StatusForbidden, gin.H{"error": "no permission to create category"})
		return
	}

	var req CategoryRequest
	// 分类的反序列化
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := model.Category{
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
	}

	if err := service.CreateCategory(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "create category success", "id": category.ID})
}

// 更新分类
func UpdateCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req CategoryRequest
	// 分类的反序列化
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := model.Category{
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
	}

	if err := service.UpdateCategory(uint(id), &category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "update category success"})
}

// 删除分类
func DeleteCategory(c *gin.Context) {
	idStr := c.Param("id")
	// 将字符串ID转换为uint
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	// 从上下文获取当前用户，检查是否为管理员
	uidVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not logged in"})
		return
	}
	uid, ok := uidVal.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user id in context"})
		return
	}

	user, err := service.GetUserByID(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user.Role != 1 {
		c.JSON(http.StatusForbidden, gin.H{"error": "no permission to delete category"})
		return
	}

	if err := service.DeleteCategory(uint(id)); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "delete category success"})
}

// 获取单个分类
func GetCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	category, err := service.GetCategoryByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

// 获取分类列表
func ListCategories(c *gin.Context) {
	categories, err := service.ListCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"list": categories})
}
