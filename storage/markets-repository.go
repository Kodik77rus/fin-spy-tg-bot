package storage

import (
	"github.com/Kodik77rus/fin-spy-tg-bot/internal/fin-spy-tg-bot/models"
)

var markets []models.Market

func (st *Storage) GetAllMarkets() ([]models.Market, error) {
	result := st.db.Find(&markets)
	if result.Error != nil {
		return nil, result.Error
	}
	return markets, nil
}
