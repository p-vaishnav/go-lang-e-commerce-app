package database

import (
	"database/sql"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDB *gorm.DB
var SqlDB *sql.DB

func InitDB() (*gorm.DB, *sql.DB) {
	var dsn string // data source name
	var err error

	dsn = " host=" + " user=" + " passwaord=" + " dbname=" + " port=" + " sslmode=disable"

	PostgresDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error while opening connection")
	}

	SqlDB, err = PostgresDB.DB()
	if err != nil {
		log.Fatal("") // ??
	}

	return PostgresDB, SqlDB
}
