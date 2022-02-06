package storage

import (
	"github.com/Kodik77rus/fin-spy-tg-bot/internal/fin-spy-tg-bot/models"
)

const pageSize = 2

var market *[]string
var markets []models.Market

type MarketResponse struct {
	Markets []models.Market
	Count   int
}

func (st *Storage) GetAllMarkets(page int) (*MarketResponse, error) {
	result := st.db.
		Select(
			"name",
			"code",
			"mic",
			"location",
			"country",
			"city",
			"delay",
			"hour",
		).
		Limit(pageSize).
		Offset(pageSize * (page - 1)).
		Find(&markets)
	if result.Error != nil {
		return nil, result.Error
	}

	response := MarketResponse{
		Markets: markets,
		Count:   int(result.RowsAffected),
	}
	return &response, nil
}

func (st *Storage) FindMarketsWithParam(param string) (*[]string, error) {
	result := st.db.
		Model(&models.Market{}).
		Distinct().
		Order(param+" asc").
		Pluck(param, &market)
	if result.Error != nil {
		return nil, result.Error
	}
	return market, nil
}