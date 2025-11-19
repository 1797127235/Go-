package v1

import (
	"blog-server/service"
	"blog-server/utils"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type RegisterReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var req RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.CreateUser(req.Username, req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "register success"})
}

func Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := service.CheckUser(req.Username, req.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "密码错误请重试"})
		return
	}

	token, err := utils.GenerateToken(u.ID, u.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":    token,
		"user_id":  u.ID,
		"username": u.Username,
		"avatar": u.Avatar,
	})
}

// 上传头像接口
func UploadAvatar(c *gin.Context) {
	// 从上下文中获取用户ID（JWT 中间件写入的）
	uid, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}
	userID, ok := uid.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户ID类型错误"})
		return
	}

	//获取上传的文件
	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//生成唯一文件名  时间戳 + 原文件后缀（带上用户ID，避免冲突）
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%d_%d%s", userID, time.Now().Unix(), ext)

	//保存文件
	if err := c.SaveUploadedFile(file, "./images/"+filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//更新用户头像
	fileUrl := "/images/" + filename
	fileUrl = "http://localhost:8080" + fileUrl
	fmt.Println(fileUrl)
	if err := service.UpdateUserAvatar(userID, fileUrl); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//返回文件的URL给前端
	c.JSON(http.StatusOK, gin.H{
		"url": fileUrl,
		"msg": "上传头像成功",
	})
}
