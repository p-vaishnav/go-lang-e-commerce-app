package database

import (
	"backend-commerce/configs"
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDB *gorm.DB
var SqlDB *sql.DB

func InitDB() (*gorm.DB, *sql.DB) {
	var dsn string // data source name a.k.a connection string
	var err error

	dsn = " host=" + configs.DB.Host +
		" user=" + configs.DB.Username +
		" password=" + configs.DB.Password +
		" dbname=" + configs.DB.Database +
		" port=" + configs.DB.Port +
		" sslmode=disable"

	// postgres.Open calls the psql driver
	// establishing connection to psql using the GORM library
	PostgresDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error while opening connection")
	}

	// it gives sql.DB which helps for more control over the database
	SqlDB, err = PostgresDB.DB()
	if err != nil {
		log.Fatal("") // ??
	}

	fmt.Println("connected to database successfully")
	return PostgresDB, SqlDB
}
