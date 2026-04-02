package sqlite

import (
	"github.com/kakebon/backend/domain/model"
	"gorm.io/gorm"
)

type categoryPersistence struct {
	DB *gorm.DB	
}

func NewCategoryPersistence(db *gorm.DB) *categoryPersistence {
	return &categoryPersistence{}
}

func (p *categoryPersistence) GetAll () ([]model.Category, error) {
	var categories []model.Category
	if err := p.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}