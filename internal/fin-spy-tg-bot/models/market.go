package models

type Market struct {
	Name      string `json:"name"`
	Code      string `json:"code"`
	Mic       string `json:"mic" gorm:"primaryKey"`
	YahooCode string `json:"yahooCode"`
	Location  string `json:"location"`
	Country   string `json:"country"`
	City      string `json:"city"`
	Delay     uint   `json:"delay"`
	Hour      string `json:"hour"`
}
