package storage

import (
	"github.com/Kodik77rus/fin-spy-tg-bot/internal/fin-spy-tg-bot/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Instance of storage
type Storage struct {
	config *Config
	db     *gorm.DB //orm instance
}

//set url from app.config into storage config
func New(url string, config *Config) *Storage {
	config.DatabaseURL = url
	return &Storage{
		config: config,
	}
}

//connection method for Db
func (storage *Storage) Open() error {
	db, err := gorm.Open(postgres.Open(storage.config.DatabaseURL), &gorm.Config{})
	if err != nil {
		return err
	}
	db.AutoMigrate(&models.User{}, &models.Watchlist{}, &models.Asset{}, &models.Market{})
	storage.db = db
	return nil
}

//method for close db connection
func (storage *Storage) Close() error {
	db, err := storage.db.DB()
	if err != nil {
		return err
	}

	if err := db.Ping(); err == nil {
		if db.Close(); err != nil {
			return err
		}
	}

	return nil
}
