package storage

//storage config
type Config struct {
	DatabaseURL string // Строка подключения к БД
}

//create base  storage config
func NewConfig() *Config {
	return &Config{}
}
