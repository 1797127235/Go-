package router

import(
	v1 "blog-server/api/v1"

	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine{
	r := gin.Default()
	
	api := r.Group("/api")
	{
		api.POST("/user/register",v1.Register)
		api.POST("/user/login",v1.Login)
	}

	return r
}
