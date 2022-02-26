package app

import (
	"os"

	"github.com/Kodik77rus/fin-spy-tg-bot/storage"
)

//Default config instance of app
type Config struct {
	TgBot       string //Tg Token
	LoggerLevel string
	DatabaseURL string
	FinhubToken string
	Storage     *storage.Config //DB instance
}

//Create default app config
func NewConfig() *Config {
	return &Config{
		TgBot:       getEnv("TG_BOT", ""),
		Storage:     storage.NewConfig(),
		LoggerLevel: getEnv("LOG_LVL", "debug"),
		DatabaseURL: getEnv("DB_URL", ""),
		FinhubToken: getEnv("FIN_HUB", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
