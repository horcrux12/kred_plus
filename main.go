package main

import (
	"fmt"
	migrate "github.com/rubenv/sql-migrate"
	"kredi-plus.com/be/app"
	"kredi-plus.com/be/config"
	"log"
)

func main() {
	config.GenerateConfiguration()
	app.InitApplicationAttribute()
	DBMigration()
}

func DBMigration() {
	db, err := app.KrediApp.DBConn.DB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	migrations := &migrate.FileMigrationSource{Dir: "db/migrations"}
	n, err := migrate.Exec(db, "mysql", migrations, migrate.Up)
	if err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	fmt.Printf("Applied %d migrations!\n", n)
}
