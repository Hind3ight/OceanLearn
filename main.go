package main

import (
	"fmt"
	"github.com/Hind3ight/OceanLearn/pkg/lib"
	"github.com/Hind3ight/OceanLearn/pkg/model"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	db := InitDB()
	r := gin.Default()
	r.POST("/api/auth/register", func(c *gin.Context) {
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
		db.Create(&user)

		c.JSON(200, gin.H{
			"message": "注册成功",
			"name":    name,
		})
	})
	r.Run(":8082") // 监听并在 0.0.0.0:8080 上启动服务
}

func InitDB() *gorm.DB {
	host := "localhost"
	port := 3306
	database := "ginDB"
	username := "root"
	password := "123456"
	charset := "utf8"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset)
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database , err:" + err.Error())
	}
	user := model.User{}
	err = db.AutoMigrate(&user)
	return db
}
