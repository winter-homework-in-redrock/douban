package middleware

import (
	"douban/tool"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//JWTAuthMiddleware JWT中间件，参照李文周博客
func JWTAuthMiddleware() func(*gin.Context) {
	return func(c *gin.Context) {
		//客户端携带Token放在Header的Authorization中，并使用Bearer开头
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"status": "-1",
				"error":  "authHeader is empty",
			})
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(200, gin.H{
				"status": "-2",
				"error":  "authHeader is wrong",
			})
			c.Abort()
			return
		}
		//解析token
		claims, err := tool.ParseToken(parts[1])
		if err != nil {
			c.JSON(200, gin.H{
				"status": "-3",
				"error":  "invalid Token",
			})
			c.Abort()
			return
		}
		c.Set("phone", claims.Phone)
		c.Next()
	}
}
