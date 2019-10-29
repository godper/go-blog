package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors 跨域配置
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie", "token"}
	config.AllowOrigins = []string{"http://localhost:8080", "http://www.godper.com", "http://192.168.0.101:8080", "http://192.168.0.101:8081"}
	config.AllowCredentials = true
	return cors.New(config)
}
