package models

import (
	"github.com/lib/pq"
)

type User struct {
	Id       uint           `json:"id" gorm:"primaryKey"`
	UserName string         `json:"first_name" gorm:"NOT NULL"`
	Language string         `json:"language_code" gorm:"NOT NULL"`
	Markets  pq.StringArray `json:"markets" gorm:"type:text[]" `
}
