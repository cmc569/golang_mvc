package main

import (
	"api/config/setting"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	err := godotenv.Load(".env")
	if err == nil {
		gin.SetMode(gin.DebugMode)
		setting.InitConfig()
	}
	fmt.Println(gin.Mode())
	// 啟動mysql_db長連接
	//db_gorm.OpenConnect()
	fmt.Println(setting.ServerConfig.Port)
	router := Router()
	srv := &http.Server{
		Addr:    ":" + setting.ServerConfig.Port,
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	log.Println("Server exiting")
}
