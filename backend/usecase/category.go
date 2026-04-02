package usecase

import (
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