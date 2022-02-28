package middleware

import (
	"github.com/Hind3ight/OceanLearn/common"
	"github.com/Hind3ight/OceanLearn/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "请求头中auth为空",
			})
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "请求头中token格式错误",
			})
			c.Abort()
			return
		}

		mc, err := common.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "无效的token",
			})
			c.Abort()
			return
		}
		userId := mc.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		if user.ID == 0 {
			c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
				"code": 203,
				"msg":  "用户权限不足",
			})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}

}
