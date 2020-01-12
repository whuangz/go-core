package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("../src/config/config.env"); err != nil {
		log.Print("No .env file found")
	}
}

func GetConfig(name string) string {
	envData, ok := os.LookupEnv(name)
	if !ok {
		panic(fmt.Sprintf("%s env variable not set", name))
	}

	return envData
}
