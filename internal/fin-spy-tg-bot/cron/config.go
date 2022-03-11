package cron

import (
	"time"

	"github.com/Kodik77rus/fin-spy-tg-bot/storage"
	"github.com/sirupsen/logrus"

	"github.com/go-resty/resty/v2"
	"github.com/robfig/cron/v3"
)

type Cron struct {
	cron         *cron.Cron
	finHubClient *resty.Client
	yahooClient  *resty.Client
	storage      *storage.Storage
	logger       *logrus.Logger
}

type finHubResponse struct {
	CompanyName string `json:"symbol"`
	// CompanyType string `json:"type"`
}

func New(finhubToken string, st *storage.Storage, logger *logrus.Logger) *Cron {
	moscow, _ := time.LoadLocation("Europe/Moscow")

	return &Cron{
		cron:         cron.New(cron.WithLocation(moscow)),
		finHubClient: finHubClient(finhubToken),
		yahooClient:  yahooClient(),
		storage:      st,
		logger:       logger,
	}
}

func finHubClient(token string) *resty.Client {
	client := resty.New()

	client.SetBaseURL("https://finnhub.io/api/v1")
	client.SetHeader("X-Finnhub-Token", token)

	return client
}

func yahooClient() *resty.Client {
	client := resty.New()

	client.BaseURL = ""

	return client
}
