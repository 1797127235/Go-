package v1

import (
	"blog-server/model"
	"blog-server/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentRequest struct {
	Content  string `json:"content" binding:"required"` //评论内容
	ParentID uint   `json:"parent_id"`                  //父评论ID
}

// 创建评论（针对某篇文章）
func CreateComment(c *gin.Context) {
	// 从路径中获取文章 ID
	idStr := c.Param("id")
	postID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	/*
		从中间件注入的上下文中取信息
	*/
	uidVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not logged in"})
		return
	}
	
	/*
		类型断言
		从中间件取出来的东西是interface{}类型，需要断言成uint类型
	*/
	uid, ok := uidVal.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user id in context"})
		return
	}

	var req CommentRequest
	// 反序列化请求体
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment := model.Comment{
		UserID:   uid,
		PostID:   uint(postID),
		Content:  req.Content,
		ParentID: req.ParentID,
	}

	// 创建评论
	if err := service.CreateComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "create comment success",
		"id":      comment.ID,
	})
}

// 查看某篇文章下的评论列表
func ListComments(c *gin.Context) {
	idStr := c.Param("id")
	postID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	comments, err := service.ListComments(uint(postID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"list": comments})
}
