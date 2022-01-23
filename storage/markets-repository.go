package storage

import (
	"github.com/Kodik77rus/fin-spy-tg-bot/internal/fin-spy-tg-bot/models"
)

const pageSize = 2

var markets []models.Market

func (st *Storage) GetAllMarkets(page int) ([]models.Market, error) {
	result := st.db.Limit(pageSize).Offset(pageSize * (page - 1)).Find(&markets)
	if result.Error != nil {
		return nil, result.Error
	}
	return markets, nil
}
