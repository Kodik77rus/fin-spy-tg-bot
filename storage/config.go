package storage

//storage config
type Config struct {
	// Строка подключения к БД
	DatabaseURL string
}

//create base  storage config
func NewConfig() *Config {
	return &Config{}
}
