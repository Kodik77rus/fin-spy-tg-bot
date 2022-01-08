package app

import "github.com/Kodik77rus/fin-spy-tg-bot/storage"

//Default config instance of app server
type Config struct {
	TgBot       string          `toml:"TG_BOT"` //Tg Token
	LoggerLevel string          `toml:"LOGGER_LVL"`
	DatabaseURL string          `toml:"DATABASE_URL"`
	Storage     *storage.Config //DB instance
}

//Create default app config
func NewConfig() *Config {
	return &Config{
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}
