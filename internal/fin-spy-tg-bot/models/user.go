package models

type User struct {
	Watchlist []Watchlist
<<<<<<< HEAD
	Id        uint   `json:"id" gorm:"primaryKey"`
	UserName  string `json:"first_name"`
=======
	Id        uint   `json: "id;"gorm: "primaryKey"`
	UserName  string `json: "first_name"`
>>>>>>> b7608121ac40f6ee6acc1f4b47dcbe84db968401
	Language  string `json:"language_code"`
}
