package configs

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// NOTE: split_words should be inside quotes
type DBConfig struct {
	Host     string `split_words:"true" json:"DB_HOST"`
	Port     string `split_words:"true" json:"DB_PORT"`
	Username string `split_words:"true" json:"DB_USERNAME"`
	Database string `split_words:"true" json:"DB_DATABASE"`
	Password string `split_words:"true" json:"DB_PASSWORD"`
}

var DB *DBConfig

func loadDBConfig() {
	DB = &DBConfig{}
	err := envconfig.Process("db", DB)

	if err != nil {
		log.Fatal(err.Error())
	}
}
