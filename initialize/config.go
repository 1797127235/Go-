package initialize

import (
    "blog-server/config"
    "blog-server/global"
    "log"
    "os"

    "gopkg.in/yaml.v3"
)

func InitConfig() {
    data, err := os.ReadFile("config.yaml")
    if err != nil {
        log.Fatal("read config.yaml failed:", err)
    }

    c := &config.ServerConfig{}
    if err := yaml.Unmarshal(data, c); err != nil {
        log.Fatal("parse config.yaml failed:", err)
    }

    global.Config = c
}
