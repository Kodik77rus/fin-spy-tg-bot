package models

import "gorm.io/gorm"

type Asset struct {
	gorm.Model
<<<<<<< HEAD
	Price          uint   `json:"regularMarketPrice"`
	Ticker         string `json:"ticker"`
	ShortName      string `json:"shortName"`
	LongName       string `json:"longName"`
	Exchange       string `json:"exchange"`
	ExchangeSymbol string `json:"exchangeSymbol"`
	Currency       string `json:"currency"`
	Country        string `json:"country"`
	City           string `json:"city"`
	Sector         string `json:"sector"`
	Industry       string `json:"industry"`
	QuoteType      string `json:"quoteType"`
	Website        string `json:"website"`
	IsWatch        bool   `json:"isWatch"`
=======
	Price          uint   `json :"regularMarketPrice"`
	Ticker         string `json :"ticker"`
	ShortName      string `json :"shortName"`
	LongName       string `json :"longName"`
	Exchange       string `json :"exchange"`
	ExchangeSymbol string `json :"exchangeSymbol"`
	Currency       string `json :"currency"`
	Country        string `json :"country"`
	City           string `json :"city"`
	Sector         string `json :"sector"`
	Industry       string `json :"industry"`
	QuoteType      string `json :"quoteType"`
	Website        string `json :"website"`
	IsWatch        bool   `json :"IsWatch"`
>>>>>>> b7608121ac40f6ee6acc1f4b47dcbe84db968401
}
