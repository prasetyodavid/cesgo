package config

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := "sqlserver://sa:1234@localhost:53726?database=gl_2018"
	var err error

	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Error),
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("Cannot connect database")
	} else {
		fmt.Println("Connected")
	}
	DB.AutoMigrate()
}
