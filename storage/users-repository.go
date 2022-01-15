package storage

import (
	"github.com/Kodik77rus/fin-spy-tg-bot/internal/fin-spy-tg-bot/models"
	"gorm.io/gorm"
)

var user *models.User

func (st *Storage) CreateUser(user *models.User) error {
	if err := st.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (st *Storage) FindUser(id uint) (*gorm.DB, error) {
	result := st.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

func (st *Storage) UpdateUser(user *models.User) error {
	if err := st.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}
