package app

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kredi-plus.com/be/config"
	"time"
)

func newMysqlConnect() (DB *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Attr.Mysql.Username,
		config.Attr.Mysql.Password,
		config.Attr.Mysql.Host,
		config.Attr.Mysql.Port,
		config.Attr.Mysql.Database)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	// Dapatkan sql.DB untuk setting pool
	sqlDB, err := database.DB()
	if err != nil {
		return
	}

	// Setting connection pooling
	sqlDB.SetMaxIdleConns(10)                  // Maksimal idle connections di pool
	sqlDB.SetMaxOpenConns(100)                 // Maksimal total open connections ke database
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // Maksimal waktu hidup connection

	DB = database
	fmt.Println("Database connection successfully opened and configured!")
	return
}
