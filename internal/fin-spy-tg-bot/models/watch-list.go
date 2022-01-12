package models

type Watchlist struct {
	User   User
<<<<<<< HEAD
	UserId uint   `json:"UserId"`
	Price  uint   `json:"regularMarketPrice"`
	Ticker string `json:"ticker"`
=======
	UserId uint   `json :"UserId"`
	Price  uint   `json :"regularMarketPrice"`
	Ticker string `json :"ticker"`
>>>>>>> b7608121ac40f6ee6acc1f4b47dcbe84db968401
}
