package setting

import (
	"github.com/gin-gonic/gin"
	"os"
)

type configPath struct {
	ErrorLogPath      string
	HttpDirectoryRoot string
}

var PathConfig configPath

func pathConfig() {
	switch gin.Mode() {
	case gin.ReleaseMode:
		PathConfig = configPath{
			"/var/log/error.log",
			"https://img.file-server.com"}
	case gin.DebugMode:
		PathConfig = configPath{
			os.Getenv("ERROR_LOG_PATH"),
			os.Getenv("FILE_URL")}
	case gin.TestMode:
		PathConfig = configPath{
			"/var/log/error.log",
			"https://img.file-server.com"}
	}
}
