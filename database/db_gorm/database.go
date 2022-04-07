package db_gorm

import (
	"api/config/setting"
	"api/model"
	"api/util/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func OpenConnect() {
	var err error

	dsn := setting.DataBaseConfig.UserName + ":" + setting.DataBaseConfig.Password + "@tcp(" + setting.DataBaseConfig.Host + ":" + setting.DataBaseConfig.Port + ")/" + setting.DataBaseConfig.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Error(err)
	}

	if err != nil {
		log.Error(err)
	}

	sqlDB, _ := db.DB()
	//設置閒置連接
	sqlDB.SetMaxIdleConns(10)
	//設置最多併發數
	sqlDB.SetMaxOpenConns(100)

	model.SetDB(db)
}


func GetDB() *gorm.DB {
	return db
}
