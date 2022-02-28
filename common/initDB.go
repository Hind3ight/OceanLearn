package common

import (
	"fmt"
	"github.com/Hind3ight/OceanLearn/pkg/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDB() *gorm.DB {
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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database , err:" + err.Error())
	}
	user := model.User{}
	err = db.AutoMigrate(&user)
	return db
}

func GetDB() *gorm.DB {
	return initDB()
}
