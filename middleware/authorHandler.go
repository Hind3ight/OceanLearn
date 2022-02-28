package middleware

import (
	"github.com/Hind3ight/OceanLearn/common"
	"github.com/Hind3ight/OceanLearn/model"
	"github.com/Hind3ight/OceanLearn/pkg/lib"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			lib.Response(c, http.StatusUnprocessableEntity, 422, nil, "请求头中auth为空")
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			lib.Response(c, http.StatusUnprocessableEntity, 422, nil, "请求头中token格式错误")
			c.Abort()
			return
		}

		mc, err := common.ParseToken(parts[1])
		if err != nil {
			lib.Response(c, http.StatusUnprocessableEntity, 422, nil, "无效的token")
			c.Abort()
			return
		}
		userId := mc.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		if user.ID == 0 {
			lib.Response(c, http.StatusNonAuthoritativeInfo, 203, nil, "用户权限不足")
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}

}
