package model

import (
	"time"
)

type Users struct {
	Id          int64  `gorm:"type:bigint(20) not null auto_increment; primary_key;"`
	LineToken   string `gorm:"type:varchar(50) not null"`
	Picture     string `gorm:"type:varchar(255) not null"`
	Name        string `gorm:"type:varchar(20) not null"`
	Phone       string `gorm:"type:varchar(8) not null"`
	CountryCode string `gorm:"type:varchar(5) not null"`
	Status      int64  `gorm:"type:int(1) not null; comment: 1:未開通, 2:開通; default:1"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
