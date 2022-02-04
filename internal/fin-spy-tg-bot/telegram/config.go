package telegram

type Config struct {
	RuDictionary RuDictionary
	EnDictionary EnDictionary
}

type RuDictionary struct {
	startMessage string
	setLanguage  string
	RuEror       RuEror
}

type EnDictionary struct {
	startMessage string
	setLanguage  string
	EnEror       EnEror
}

type RuEror struct {
}

type EnEror struct {
}

func NewConfig() *Config {
	return &Config{
		RuDictionary: RuDictionary{
			startMessage: "Скоро бот начнёт работать",
			setLanguage:  "Вы выбрали русский язык!",
		},
		EnDictionary: EnDictionary{
			startMessage: "The bot will start working soon",
			setLanguage:  "You have chosen English!",
		},
	}
}
