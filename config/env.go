package config

import (
	"fmt"
	"os"
)

type Config struct{
	PublicHost string
	Port string

	DBUser string
	DBPassword string
	DBAddress string
	DBName string
}

//Just a global variable that holds all the env variable
//We are using it so that we don't always reinitialise the config function.
var Envs = initConfig();

func initConfig() Config{
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port: getEnv("PORT", "8080"),
		DBUser: getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "mypassword"),
		DBAddress: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName: getEnv("DB_NAME", "goCommerce"),
	}
}

func getEnv(key, fallback string) string{
	if value, ok := os.LookupEnv(key); ok{
		return value;
	}

	return fallback;
} 