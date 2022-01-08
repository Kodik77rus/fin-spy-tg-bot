package storage

import (
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
