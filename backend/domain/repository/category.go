package repository

import "github.com/kakebon/backend/domain/model"

type CategoryRepository interface {
	GetAll() ([]model.Category, error)
	Create(category *model.Category) error
}