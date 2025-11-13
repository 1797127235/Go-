package initialize

import (
    "github.com/gin-gonic/gin"
    "blog-server/router"
)

func InitRouter() *gin.Engine {
    return router.SetupRouter()
}
