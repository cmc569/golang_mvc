package main

import (
	"api/config/setting"
	"api/database/db_gorm"
	"api/util/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err == nil {
		gin.SetMode(gin.DebugMode)
		setting.InitConfig()
	}
	fmt.Println(gin.Mode())
	// 啟動mysql_db長連接
	db_gorm.OpenConnect()
	fmt.Println(setting.ServerConfig.Port)
	log.Error(Router().Run(":" + setting.ServerConfig.Port))

}


