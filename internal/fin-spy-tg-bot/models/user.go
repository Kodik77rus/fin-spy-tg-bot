package models

type User struct {
	Watchlist []Watchlist
	Id        uint   `json:"id" gorm:"primaryKey"`
	UserName  string `json:"first_name" gorm:"NOT NULL"`
	Language  string `json:"language_code" gorm:"NOT NULL"`
}
