package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DBName          string
	SSLMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func LoadDatabaseConfig() (*DatabaseConfig, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("env файл не найден")
	}
	return &DatabaseConfig{
		Host:            getEnv("DB_HOST", "localhost"),
		Port:            getEnv("DB_PORT", "localhost"),
		User:            getEnv("DB_USER", "localhost"),
		Password:        getEnv("DB_PASSWORD", "localhost"),
		DBName:          getEnv("DB_NAME", "localhost"),
		SSLMode:         getEnv("DB_SSLMODE", "localhost"),
		MaxOpenConns:    getIntEnv("DB_MAX_OPEN_CONNS", 10),
		MaxIdleConns:    getIntEnv("DB_MAX_IDLE_CONNS", 5),
		ConnMaxLifetime: time.Duration(getIntEnv("DB_CONN_MAX_LIFE_MINUTES", 30)) * time.Minute,
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getIntEnv(key string, defaultValue int) int {
	value := getEnv(key, fmt.Sprintf("%d", defaultValue))
	if intValue, err := strconv.Atoi(value); err == nil {
		return intValue
	}
	return defaultValue
}
