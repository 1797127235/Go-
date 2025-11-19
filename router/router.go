package router

import (
	v1 "blog-server/api/v1"
	"blog-server/middleware"
	"blog-server/server"

	"github.com/gin-gonic/gin"
)

/*
挂载路由
*/
func SetupRouter() *gin.Engine {
	r := gin.Default()
	//设置静态文件目录
	r.Static("/images", "./images")
	// // 加载 HTML 模板
	// r.LoadHTMLGlob("templates/*")

	api := r.Group("/api")
	{
		api.POST("/user/register", v1.Register)
		api.POST("/user/login", v1.Login)
	}

	// 需要登录的路由
	auth := api.Group("")
	auth.Use(middleware.JWTAuth())
	{
		// 文章相关路由（写操作）
		auth.POST("/posts", v1.CreatePost)
		auth.PUT("/posts/:id", v1.UpdatePost)
		auth.DELETE("/posts/:id", v1.DeletePost)

		// 评论相关路由（写操作）
		auth.POST("/posts/:id/comments", v1.CreateComment)
		auth.DELETE("/comments/:id", v1.DeleteComment)

		// 分类相关路由（写操作）
		auth.POST("/categories", v1.CreateCategory)
		auth.PUT("/categories/:id", v1.UpdateCategory)
		auth.DELETE("/categories/:id", v1.DeleteCategory)

		// 用户相关路由（写操作）
		auth.POST("/user/avatar", v1.UploadAvatar)
	}

	// 不需要登录的读操作
	api.GET("/posts/:id", v1.GetPost)
	api.GET("/posts", v1.ListPosts)
	api.GET("/posts/:id/comments", v1.ListComments)
	api.GET("/categories/:id", v1.GetCategory)
	api.GET("/categories", v1.ListCategories)

	// 前台网页路由
	r.GET("/", server.HomePage)
	r.GET("/posts/:id", server.PostDetail)

	return r
}
