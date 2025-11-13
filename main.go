package main

import (
    "blog-server/global"
    "blog-server/initialize"
    "fmt"
)

func main() {
    // 1. 读取配置
    initialize.InitConfig()
    // 2. 初始化数据库
    initialize.InitMysql() 
    // 3. 初始化路由
    r := initialize.InitRouter()

    addr := fmt.Sprintf(":%d", global.Config.Server.Port)
    if err := r.Run(addr); err != nil {
        panic(err)
    }
}
