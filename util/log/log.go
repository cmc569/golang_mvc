package log

import (
	"api/config/setting"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"runtime"
	"strings"
)


func Error(err error){
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		fmt.Printf("[ERROR] %s ,line %d: %s \n", fn, line, err.Error())
		if gin.Mode() == gin.ReleaseMode {
			var builder strings.Builder
			f, _ := os.OpenFile(setting.PathConfig.ErrorLogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			defer closeFile(f)
			log.SetOutput(f)
			_,_ = fmt.Fprintf(&builder, "[ERROR] %s ,line %d: %s \n", fn, line, err.Error())
			log.Println(builder.String())
		}
	}
}

func closeFile(f *os.File){
	if f != nil {
		_ = f.Close()
	}
}