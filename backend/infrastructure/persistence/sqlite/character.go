package sqlite

import (
	"github.com/kakebon/backend/domain/model"
	"gorm.io/gorm"
)

type characterPersistence struct {
	DB *gorm.DB
}

func NewCharacterPersistence(db *gorm.DB) *characterPersistence {
	return &characterPersistence{ DB: db }
}

func (p *characterPersistence) GetByUserId(userID string) (*model.Character, error) {
	var char model.Character
	if err := p.DB.Where("user_id = ?", userID).First(&char).Error; err != nil {
		return nil, err
	}
	return &char, nil
}

func (p *characterPersistence) Update(character *model.Character) error {
	return p.DB.Save(character).Error
}