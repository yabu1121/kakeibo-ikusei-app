package sqlite

import (
	"github.com/kakebon/backend/domain/model"
	"gorm.io/gorm"
)

type UserPersistence struct {
	DB *gorm.DB
}

func NewUserPersistence (db *gorm.DB) *UserPersistence {
	return &UserPersistence{DB: db}
}

func (p *UserPersistence) Create (user *model.User) error {
	return p.DB.Create(user).Error
}

func (p *UserPersistence) GetByEmail (email string) (*model.User, error) {
	var user model.User
	if err := p.DB.Where("email = ?", email).Preload("Character").First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
