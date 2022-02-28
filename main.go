package main

import (
	"fmt"
	fileUtils "github.com/Hind3ight/OceanLearn/pkg/lib/file"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"path"
)

func main() {
	initConfig()
	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("Server.port")
	if port != "" {
		r.Run(":" + port)
	}
	r.Run()
}

func initConfig() {
	workPath, _ := os.Getwd()
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	confDir := workPath + string(os.PathSeparator) + "conf"
	viper.AddConfigPath(confDir)
	err := viper.ReadInConfig()
	if err != nil {
		defaultConfigContent := "Server:\n  port: 8802\nMysql:\n  host: \"localhost\"\n  port: 3306\n  database: \"xxx\"\n  username: \"root\"\n  password: \"123456\"\n  charset: \"utf8\""
		fileUtils.WriteFile(path.Join(confDir, "app.yaml"), defaultConfigContent)
		fmt.Println("Default configFile has  been created! Please modify as needed!")
	}
}
