package sqlite

import (
	"github.com/google/uuid"
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
	err := p.DB.Where("user_id = ?", userID).Attrs(model.Character{
		ID:             uuid.New().String(),
		UserID:         userID,
		CurrentLevel:   1,
		CurrentExp:     0,
		ExpToNextLevel: 100,
	}).FirstOrCreate(&char).Error
	if err != nil {
		return nil, err
	}
	return &char, nil
}

func (p *characterPersistence) Update(character *model.Character) error {
	return p.DB.Save(character).Error
}