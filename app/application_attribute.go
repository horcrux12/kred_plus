package app

import (
	"gorm.io/gorm"
	"log"
)

type applicationAttr struct {
	DBConn *gorm.DB
}

var KrediApp = applicationAttr{}

func InitApplicationAttribute() {
	var err error
	KrediApp.DBConn, err = newMysqlConnect()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	KrediApp.DBConn = KrediApp.DBConn.Debug()
}
