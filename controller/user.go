package controller

import (
	"github.com/Hind3ight/OceanLearn/common"
	"github.com/Hind3ight/OceanLearn/model"
	"github.com/Hind3ight/OceanLearn/pkg/lib"
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
		lib.Response(c, http.StatusUnprocessableEntity, 402, nil, "手机号必须为11位")
		return
	}

	if len(password) < 6 {
		lib.Response(c, http.StatusUnprocessableEntity, 402, nil, "密码需大于6位")
		return
	}

	if name == "" {
		name = lib.RandomString(10)
	}

	if isTelephoneExist(DB, telephone) {
		lib.Response(c, http.StatusUnprocessableEntity, 402, nil, "手机号已存在")
		return
	} else {
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			lib.Response(c, http.StatusInternalServerError, 500, nil, "加密错误")
			return
		}

		newUser := model.User{
			Name:      name,
			Telephone: telephone,
			Password:  string(hash),
		}
		DB.Create(&newUser)
	}

	lib.Success(c, nil, "注册成功")
}

func Login(c *gin.Context) {
	DB := common.GetDB()
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	if len(telephone) != 11 {
		lib.Response(c, http.StatusUnprocessableEntity, 402, nil, "手机号必须为11位")
		return
	}

	if len(password) < 6 {
		lib.Response(c, http.StatusUnprocessableEntity, 402, nil, "密码需大于6位")
		return
	}

	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		lib.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		lib.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码错误")
		return
	}

	token, err := common.GenToken(user)
	if err != nil {
		lib.Response(c, http.StatusInternalServerError, 500, nil, "token生成失败")
		log.Printf("token generate error:%s ", err)
		return
	}

	lib.Success(c, gin.H{"token": token}, "登录成功")
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")

	lib.Response(c, http.StatusOK, 200, gin.H{"user": model.ToUserResp(user.(model.User))}, "success")
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
