package database

// difference between -t and -u flag?
import (
	"database/sql"
	"log"
	"migrations/configs"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GormDB *gorm.DB
var SqlDB *sql.DB

func Connection() (*gorm.DB, *sql.DB) {
	var dsn string

	dsn = "host=" + configs.DB.Host +
		" user=" + configs.DB.Username +
		" password=" + configs.DB.Password +
		" dbname= " + configs.DB.Database +
		" port=" + configs.DB.Port +
		" sslmode=disable"

	GormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("[databse-Connetion], Error while opening db")
	}

	SqlDB, err := GormDB.DB()
	if err != nil {
		log.Fatal("[database-DB], Error in setting postgres DB")
	}

	SqlDB.SetMaxIdleConns(10)           //  DB.SetMaxIdleConns changes the limit on the maximum number of idle connections sql.DB maintains.
	SqlDB.SetMaxOpenConns(100)          //  DB.SetMaxOpenConns imposes a limit on the number of open connections. Past this limit, new database operations will wait for an existing operation to finish, at which time sql.DB will create another connection. By default, sql.DB creates a new connection any time all the existing connections are in use when a connection is needed.
	SqlDB.SetConnMaxLifetime(time.Hour) //  DB.SetConnMaxLifetime sets the maximum length of time a connection can be held open before it is closed.

	return GormDB, SqlDB
}
