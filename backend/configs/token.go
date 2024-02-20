package configs

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type TokenConfigs struct {
	AccessSecret      string `split_word:"true" json:"TOKEN_ACCESS_SECRET"`
	AccessExpiryTime  int    `split_word:"true" json:"TOKEN_ACCESS_EXPIRY_TIME"`
	RefreshSecret     string `split_word:"true" json:"TOKEN_REFRESH_SECRET"`
	RefreshExpiryTime int    `split_word:"true" json:"TOKEN_REFRESH_EXPIRY_TIME"`
}

var Token *TokenConfigs

func loadTokenConfig() {
	Token = &TokenConfigs{}

	err := envconfig.Process("token", Token)
	if err != nil {
		log.Fatal(err.Error())
	}
}
