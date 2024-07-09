package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetConfig() *Configurations {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err, ", environtment will get from os env")
	}
	return &Configurations{
		DB_URI:     os.Getenv("DB_URI"),
		DB_DIALECT: os.Getenv("DB_DIALECT"),
		HOST:       os.Getenv("HOST"),
		PORT:       os.Getenv("PORT"),
		API_KEY:    os.Getenv("API_KEY"),
	}
}
