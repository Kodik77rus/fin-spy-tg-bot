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

// Limit(1).Find(&, name)

func (st *Storage) FindMarketsWithParams(query string, data string, page int) (*MarketResponse, error) {
	var response MarketResponse

	switch query {
	case "info":
		result := st.db.
			Where(&models.Market{
				Name: data,
			}).
			Find(&markets)
		if result.Error != nil {
			return nil, result.Error
		}

		response.Markets = markets
		response.Count = int(result.RowsAffected)

	case "location":
		result := st.db.
			Where(&models.Market{
				Location: data,
			}).
			Limit(pageSize).
			Offset(pageSize * (page - 1)).
			Find(&markets)
		if result.Error != nil {
			return nil, result.Error
		}

		response.Markets = markets
		response.Count = int(result.RowsAffected)

	case "country":
		result := st.db.
			Where(&models.Market{
				Country: data,
			}).
			Limit(pageSize).
			Offset(pageSize * (page - 1)).
			Find(&markets)
		if result.Error != nil {
			return nil, result.Error
		}

		response.Markets = markets
		response.Count = int(result.RowsAffected)

	case "city":
		result := st.db.
			Where(&models.Market{
				City: data,
			}).
			Limit(pageSize).
			Offset(pageSize * (page - 1)).
			Find(&markets)
		if result.Error != nil {
			return nil, result.Error
		}

		response.Markets = markets
		response.Count = int(result.RowsAffected)
	}
	return &response, nil
}

func (st *Storage) SortedMarketList(param string) (*[]string, error) {
	if param == "markets" {
		param = "name"
	}

	if param == "locations" {
		param = "location"
	}

	if param == "countries" {
		param = "country"
	}

	if param == "cities" {
		param = "city"
	}

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
