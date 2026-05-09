package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	ServerPort string
	ServerHost string

	JWTSecret     string
	JWTExpiration int64

	EmailHost     string
	EmailPort     string
	EmailUser     string
	EmailPassword string
	EmailFrom     string

	FrontendUrl string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	jwtExpiretion, _ := strconv.ParseInt(getEnv("JWT_EXPIRATION", "900"), 10, 64)

	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost2"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "user"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "user_management"),
		DBSSLMode:  getEnv("DB_SSL_MODE", "disable"),

		ServerPort: getEnv("SERVER_PORT", "8080"),
		ServerHost: getEnv("SERVER_HOST", "localhost"),

		JWTSecret:     getEnv("JWT_SECRET", "exe312;s;-31aa2[2-21]"),
		JWTExpiration: jwtExpiretion,

		EmailHost:     getEnv("EMAIL_HOST", "smt.google.com"),
		EmailPort:     getEnv("EMAIL_PORT", "587"),
		EmailUser:     getEnv("EMAIL_USER", "test"),
		EmailPassword: getEnv("EMAIL_PASSWORD", "tesst"),
		EmailFrom:     getEnv("EMAIL_FROM", "exmaple@web.com"),

		FrontendUrl: getEnv("FRONTEND_URL", "localhost"),
	}
}

func getEnv(key, defaultValue string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}

	return defaultValue
}
