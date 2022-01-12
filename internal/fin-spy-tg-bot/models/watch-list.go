package models

type Watchlist struct {
	User   User
	UserId uint   `json:"UserId"`
	Price  uint   `json:"regularMarketPrice"`
	Ticker string `json:"ticker"`
}
