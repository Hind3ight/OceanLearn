package common

import (
	"fmt"
	"github.com/Hind3ight/OceanLearn/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDB() *gorm.DB {
	host := viper.GetString("Mysql.host")
	port := viper.GetInt("Mysql.port")
	database := viper.GetString("Mysql.database")
	username := viper.GetString("Mysql.username")
	password := viper.GetString("Mysql.password")
	charset := viper.GetString("Mysql.charset")
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
