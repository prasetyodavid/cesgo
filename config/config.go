package config

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	db_host := "192.168.221.54"
	db_username := "sa"
	db_password := "1234!strong"
	db_name := "cashier"
	db_port := "1433"
	dsn := "sqlserver://" + db_username + ":" + db_password + "@" + db_host + ":" + db_port + "?database=" + db_name
	var err error
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("Cannot connect database")
	} else {
		fmt.Println("Connected")
	}
	DB.AutoMigrate()
}
