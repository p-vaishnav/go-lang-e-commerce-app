package configs

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/lpernett/godotenv"
)

func LoadConfigs() {
	_, file, _, ok := runtime.Caller(0)

	// fmt.Println(pc, file, line) // program_counter, file_name, line_number NOTE: it executed two time

	if !ok {
		log.Fatal("error while calling Caller")
	}
	rootPath := filepath.Join(filepath.Dir(file), "../")

	err := godotenv.Load(rootPath + "/.env")
	if err != nil {
		log.Fatal("error while loding configs")
	}

	loadDBConfig()
	loadAppConfig()
}
