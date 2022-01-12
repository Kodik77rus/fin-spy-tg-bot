package telegram

type Config struct {
	RuDictionary RuDictionary
	EnDictionary EnDictionary
}

type RuDictionary struct {
	RuEror RuEror
}

type EnDictionary struct {
	EnEror EnEror
}

type RuEror struct {
}

type EnEror struct {
}

func NewConfig() *Config {
	return &Config{}
}
