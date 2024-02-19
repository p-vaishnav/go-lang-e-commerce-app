package configs

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// NOTE: split_words should be inside quotes
type AppConfig struct {
	Env string `split_words:"true" json:"APP_ENV"`
}

var App *AppConfig

func loadAppConfig() {
	App = &AppConfig{}
	err := envconfig.Process("app", App)

	if err != nil {
		log.Fatal(err.Error())
	}
}
