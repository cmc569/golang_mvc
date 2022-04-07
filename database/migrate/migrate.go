package migrate

import (
	"api/database/db_gorm"
	"api/model"
	"api/util/log"
)

func Migrate() {
	log.Error(db_gorm.GetDB().AutoMigrate(&model.Users{}))
}