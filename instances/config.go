package instances

import (
	"log"
	"os"
	"secret-santa/config"

	"github.com/joho/godotenv"
)

var cfg *config.Config

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		log.Fatalln("config path not specified")
	}

	newCfg, err := config.LoadConfig(path)
	if err != nil {
		log.Fatal(err)
	}

	cfg = &newCfg
}

func GetConfig() *config.Config {
	return cfg
}
