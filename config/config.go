package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("../../config/config.env"); err != nil {
		log.Print("No .env file found")
	}
}

func GetConfig(name string) map[string]string {
	conf := make(map[string]string)

	envData, ok := os.LookupEnv(name)
	if !ok {
		panic(fmt.Sprintf("%s env variable not set", name))
	}
	conf[name] = envData

	return conf
}
