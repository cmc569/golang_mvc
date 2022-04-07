package setting

import (
	"github.com/gin-gonic/gin"
	"os"
)

type configDB struct {
	Type     string
	Host     string
	Port     string
	DBName   string
	UserName string
	Password string
}


var DataBaseConfig configDB

func databaseConfig() {
	switch gin.Mode() {
	case gin.ReleaseMode:
		DataBaseConfig = configDB{
			"mysql",
			"127.0.0.1",
			"3306",
			"test",
			"test",
			"password"}

	case gin.DebugMode:
		DataBaseConfig = configDB{
			os.Getenv("DB_TYPE"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_POST"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD")}

	case gin.TestMode:
		DataBaseConfig = configDB{
			"mysql",
			"127.0.0.1",
			"3306",
			"test",
			"test",
			"password"}
	}
}
