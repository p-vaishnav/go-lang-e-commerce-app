package main

import (
	"backend-commerce/app"
	"backend-commerce/configs"
	"backend-commerce/database"
)

func main() {
	configs.LoadConfigs()
	// TODO: initalize redis as well
	database.InitDB()
	app.InitRoutes()
}
