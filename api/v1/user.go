package v1

import (
    "blog-server/service"
    "blog-server/utils"
    "net/http"
    "github.com/gin-gonic/gin"
)

type RegisterReq struct{
	Username string `json:"username" binding:"required"` 
	Password string `json:"password" binding:"required"` 
}

type LoginReq struct{
	Username string `json:"username" binding:"required"` 
	Password string `json:"password" binding:"required"` 
}

func Register(c *gin.Context){
	var req RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	if err := service.CreateUser(req.Username,req.Password); err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK,gin.H{"message":"register success"})
}


func Login(c *gin.Context){
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	u,err := service.CheckUser(req.Username,req.Password)

	if err != nil{
		c.JSON(http.StatusUnauthorized,gin.H{"error":err.Error()})
		return
	}

	token,err := utils.GenerateToken(u.ID,u.Username)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"token":token,
		"user_id":u.ID,
		"username":u.Username,
	})
}

