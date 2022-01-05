package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rey/micro-demo/pkg/jwt"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取token
		token := c.Request.Header.Get("token")
		if token == "" || len(token) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "JWT验证失败, 请重新登录",
			})
			c.Abort()
			return
		}

		// 解析token
		j := jwt.NewJWT()
		parsedToken, err := j.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
				"msg": "JWT验证失败, 请重新登录",
			})
			c.Abort()
			return
		}

		// 解析成功后判断是否需要被刷新 ???
		//now := time.Now().Unix()
		//expiresAt := parsedToken.ExpiresAt
		//oneday := int64(86400) //  60 * 60 * 24

		//if expiresAt > now && expiresAt-now < oneday {
		//j.RefreshToken(token)
		//}

		c.Set("claims", parsedToken)
		c.Next()
	}
}
