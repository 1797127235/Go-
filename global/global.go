package global

import (
    "blog-server/config"
    "gorm.io/gorm"
)

var (
    Config *config.ServerConfig
    DB     *gorm.DB
)
