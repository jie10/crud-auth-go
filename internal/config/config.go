package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	DBHost             string
	DBPort             string
	DBUser             string
	DBPassword         string
	DBName             string
	JWTSecret          string
	JWTAccessExpiry    time.Duration
	JWTRefreshExpiry   time.Duration
	ServerPort         string
	ServerTimeout      time.Duration
	PasswordHashCost   int
	CORSAllowedOrigins []string
}

func LoadConfig() *Config {
	return &Config{
		DBHost:             getEnv("DB_HOST", "localhost"),
		DBPort:             getEnv("DB_PORT", "5432"),
		DBUser:             getEnv("DB_USER", "postgres"),
		DBPassword:         getEnv("DB_PASSWORD", "postgres"),
		DBName:             getEnv("DB_NAME", "productdb"),
		JWTSecret:          getEnv("JWT_SECRET", "very-secret-key"),
		JWTAccessExpiry:    getEnvAsDuration("JWT_ACCESS_EXPIRY", time.Hour),
		JWTRefreshExpiry:   getEnvAsDuration("JWT_REFRESH_EXPIRY", time.Hour*24*7),
		ServerPort:         getEnv("SERVER_PORT", "8080"),
		ServerTimeout:      getEnvAsDuration("SERVER_TIMEOUT", time.Second*30),
		PasswordHashCost:   getEnvAsInt("PASSWORD_HASH_COST", 12),
		CORSAllowedOrigins: getEnvAsSlice("CORS_ALLOWED_ORIGINS", []string{"*"}, ","),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	if value, exists := os.LookupEnv(key); exists {
		duration, err := time.ParseDuration(value)
		if err == nil {
			return duration
		}
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvAsSlice(key string, defaultValue []string, sep string) []string {
	if value, exists := os.LookupEnv(key); exists {
		return strings.Split(value, sep)
	}
	return defaultValue
}
