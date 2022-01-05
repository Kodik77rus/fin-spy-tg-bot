package api

import "github.com/Kodik77rus/fin-spy-tg-bot/storage"

//default config instance of API server
type Config struct {
	//Port
	Port string `toml:"PORT"`

	//For connect Tg API
	TgBot string `toml:"TG_BOT"`

	//Logger Level
	LoggerLevel string `toml:"LOGGER_LVL"`

	//
	DatabaseURL string `toml:"DATABASE_URL"`

	//
	Storage *storage.Config
}

//create default api config
func NewConfig() *Config {
	return &Config{
		Port:        ":8080",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}
