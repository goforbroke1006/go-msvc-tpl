package config

import "github.com/joho/godotenv"

func LoadFromEnvFile() error {
	return godotenv.Load("./.env")
}
