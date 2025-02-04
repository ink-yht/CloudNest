package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ink-yht/CloudNest/internal/web/user_web"
)

func main() {
	server := gin.Default()

	c := user_web.NewUserHandler()
	c.RegisterRoutes(server)
	// 你这还有别的路由
	server.Run(":8080")
}
