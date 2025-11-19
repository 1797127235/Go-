package server

import (
	"net/http"
	"strconv"

	"blog-server/service"

	"github.com/gin-gonic/gin"
)

// 博客首页：文章列表
func HomePage(c *gin.Context) {
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
		c.String(http.StatusInternalServerError, "failed to load posts: %v", err)
		return
	}

	categories, err := service.ListCategories()
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to load categories: %v", err)
		return
	}

	// 计算简单分页信息
	hasPrev := page > 1
	hasNext := int64(page*pageSize) < total
	prevPage := 1
	nextPage := page
	if hasPrev {
		prevPage = page - 1
	}
	if hasNext {
		nextPage = page + 1
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Posts":      posts,
		"Categories": categories,
		"Page":       page,
		"PageSize":   pageSize,
		"Total":      total,
		"HasPrev":    hasPrev,
		"HasNext":    hasNext,
		"PrevPage":   prevPage,
		"NextPage":   nextPage,
	})
}

// 文章详情页
func PostDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id == 0 {
		c.String(http.StatusBadRequest, "invalid post id")
		return
	}

	post, err := service.GetPostByID(uint(id))
	if err != nil {
		c.String(http.StatusNotFound, "post not found")
		return
	}

	comments, err := service.ListComments(uint(id))
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to load comments: %v", err)
		return
	}

	c.HTML(http.StatusOK, "post.html", gin.H{
		"Post":     post,
		"Comments": comments,
	})
}
