package usecase

import (
	"github.com/google/uuid"
	"github.com/kakebon/backend/domain/model"
	"github.com/kakebon/backend/domain/repository"
)

type CategoryUsecase struct {
	repo repository.CategoryRepository
}

func NewCategoryUsecase (repo repository.CategoryRepository) *CategoryUsecase {
	return &CategoryUsecase{repo: repo}
}

func (u *CategoryUsecase) GetAll() ([]model.Category, error) {
	return u.repo.GetAll()
}

func (u *CategoryUsecase) Create (name string) (*model.Category, error) {
	category := &model.Category{
		ID: uuid.New().String(),
		Name: name,
	}
	if err := u.repo.Create(category); err != nil {
		return nil, err
	}
	return category, nil
}