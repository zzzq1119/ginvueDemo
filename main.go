package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r = CollectRoute(r)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
