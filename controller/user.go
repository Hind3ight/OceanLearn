package controller

import (
	"github.com/Hind3ight/OceanLearn/common"
	"github.com/Hind3ight/OceanLearn/pkg/lib"
	"github.com/Hind3ight/OceanLearn/pkg/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
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

	if isTelephoneExist(DB, telephone) {
		c.JSON(200, gin.H{
			"message": "手机号已存在",
		})
		return
	} else {
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "加密错误",
			})
		}

		newUser := model.User{
			Name:      name,
			Telephone: telephone,
			Password:  string(hash),
		}
		DB.Create(&newUser)
	}

	c.JSON(200, gin.H{
		"message": "注册成功",
		"name":    name,
	})
}

func Login(c *gin.Context) {
	DB := common.GetDB()
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

	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户不存在",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码错误",
		})
		return
	}

	token, err := common.GenToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "token生成失败",
		})
		log.Printf("token generate error:%s ", err)
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"data":    gin.H{"token": token},
		"message": "登录成功",
	})
}

func Info(c *gin.Context) {
	username, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    gin.H{"user": username},
		"message": "success",
	})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
