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