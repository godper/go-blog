package middleware

import (
	"blog/http/response"
	"blog/service"
	ujwt "blog/util/jwt"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

// JwtAuth 中间件，检查token
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := response.NewR(c)
		srv := service.Srv

		tokenstring := c.Request.Header.Get("token")
		if tokenstring == "" {
			r.FailedResponse(errors.New("请注册并登录后访问"))
			c.Abort()
			return
		}

		// parseToken 解析token包含的信息

		claims, err := srv.Jwt.ParseToken(tokenstring)
		if err != nil {
			if err == ujwt.ErrTokenExpired {
				//授权过期 延迟时间2小时
				if claims == nil || (claims.ExpiresAt+7200) < time.Now().Unix() {
					r.TokenErrResponse(errors.New("授权已过期，请重新登录"))
					c.Abort()
					return
				}
				//更新token
				newtoken, err := srv.Jwt.RefreshToken(claims)
				if err == nil {
					c.Set("newtoken", newtoken)
				}
			} else {
				//token无效
				r.TokenErrResponse(errors.New("无效登录"))
				c.Abort()
				return
			}
		}

		// 继续交由下一个路由处理,并将解析出的信息传递下去

		if claims.Role == "user" {
			c.Set("userID", claims.ID)
		}
		if claims.Role == "admin" {
			c.Set("adminID", claims.ID)
		}

		// }
		c.Next()
	}
}
