package configs

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func LoadConfigs() {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("error while loading configs")
	}

	rootPath := filepath.Join(filepath.Dir(file), "../")
	fmt.Println("[rootPath]", rootPath)
	err := godotenv.Load(rootPath + "/.env")
	if err != nil {
		log.Println("error while loading .env")
	}

	loadOTPConfig()
}

// TODO: I can write my own .env reader file
