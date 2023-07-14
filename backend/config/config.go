package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"path/filepath"
	"vk-notification-monitor/utils"
)

type Config struct {
	SERVER_ADDRESS       string
	DB_HOST              string
	DB_PORT              string
	DB_USER              string
	DB_PASSWORD          string
	DB_NAME              string
	DB_SSL               string
	DB_DRIVER            string
	VK_TOKEN             string
	VK_VERSION           string
	VK_API_URL           string
	ENABLE_NOTIFICATIONS int
	TELEGRAM_BOT_TOKEN   string
	SYNC_INTERVAL        int
}

func NewConfig() *Config {

	envFile := filepath.Join(".", ".env")
	absPath, err := filepath.Abs(envFile)

	err = godotenv.Load(absPath)
	if err != nil {
		fmt.Println(absPath)
		log.Println(err)
	}

	cfg := Config{
		SERVER_ADDRESS:       utils.LookupEnvString("SERVER_ADDRESS", "0.0.0.0:8000"),
		DB_HOST:              utils.LookupEnvString("DB_HOST", "localhost"),
		DB_PORT:              utils.LookupEnvString("DB_PORT", "5432"),
		DB_USER:              utils.LookupEnvString("DB_USER", "AdminUser"),
		DB_NAME:              utils.LookupEnvString("DB_NAME", "monitor"),
		DB_PASSWORD:          utils.LookupEnvString("DB_PASSWORD", "AdminPassword"),
		DB_SSL:               utils.LookupEnvString("DB_SSL", "disable"),
		DB_DRIVER:            utils.LookupEnvString("DB_DRIVER", "postgres"),
		VK_TOKEN:             utils.LookupEnvString("VK_TOKEN", "VK_TOKEN"),
		VK_VERSION:           utils.LookupEnvString("VK_VERSION", "5.131"),
		VK_API_URL:           utils.LookupEnvString("VK_API_URL", "https://api.vk.com"),
		ENABLE_NOTIFICATIONS: utils.LookupEnvInt("ENABLE_NOTIFICATIONS", 0),
		TELEGRAM_BOT_TOKEN:   utils.LookupEnvString("TELEGRAM_BOT_TOKEN", ""),
		SYNC_INTERVAL:        utils.LookupEnvInt("SYNC_INTERVAL", 320),
	}

	return &cfg
}
