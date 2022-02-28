package controller

import (
	"github.com/Hind3ight/OceanLearn/common"
	"github.com/Hind3ight/OceanLearn/pkg/lib"
	"github.com/Hind3ight/OceanLearn/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	DB := common.GetDB()
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    402,
			"message": "手机号必须为11位",
		})
		return
	}

	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    402,
			"message": "密码需大于6位",
		})
		return
	}

	if name == "" {
		name = lib.RandomString(10)
	}

	user := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	DB.Create(&user)

	c.JSON(200, gin.H{
		"message": "注册成功",
		"name":    name,
	})
}
