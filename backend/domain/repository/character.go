package repository

import "github.com/kakebon/backend/domain/model"

type CharacterRepository interface {
	GetByUserId (userID string) (*model.Character, error)
	Update(character *model.Character) error
}