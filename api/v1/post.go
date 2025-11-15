package v1

/*
	文章管理的接口
*/
import (
	"errors"
	"net/http"
	"strconv"

	"blog-server/model"
	"blog-server/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostRequest struct {
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
	CategoryID uint   `json:"category_id" binding:"required"`
	Status     uint   `json:"status"`
}

// 创建文章
func CreatePost(c *gin.Context) {
	var req PostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 从上下文中获取当前登录用户ID
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

	post := model.Post{
		Title:      req.Title,
		Content:    req.Content,
		AuthorID:   uid,
		CategoryID: req.CategoryID,
		Status:     req.Status,
	}

	if post.Status == 0 {
		post.Status = 1
	}

	if err := service.CreatePost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "create post success",
		"id":      post.ID,
	})
}

// 根据文章ID获取文章
func GetPost(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	post, err := service.GetPostByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

// 分页获取文章列表
func ListPosts(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	sizeStr := c.DefaultQuery("page_size", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(sizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	posts, total, err := service.ListPosts(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"list":      posts,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// 更新文章
func UpdatePost(c *gin.Context) {
	idStr := c.Param("id") //获取路径参数ID
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	//反序列化
	var req PostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	update := model.Post{
		Title:      req.Title,
		Content:    req.Content,
		CategoryID: req.CategoryID,
		Status:     req.Status,
	}

	//更新文章
	if err := service.UpdatePost(uint(id), &update); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//返回成功信息
	c.JSON(http.StatusOK, gin.H{"message": "update post success"})
}

// 根据文章ID删除文章
func DeletePost(c *gin.Context) {
	idStr := c.Param("id")
	//将字符串ID转换为uint
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	//根据ID删除文章
	if err := service.DeletePost(uint(id)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "delete post success"})
}
