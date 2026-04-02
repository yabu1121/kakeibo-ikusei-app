package usecase

import (
	"github.com/kakebon/backend/domain/model"
	"github.com/kakebon/backend/domain/repository"
	"github.com/kakebon/backend/domain/service"
)

type CharacterUsecase struct {
	repo repository.CharacterRepository
}

func NewCharacterUsecase(r repository.CharacterRepository) *CharacterUsecase {
	return &CharacterUsecase{repo: r}
}

func (u *CharacterUsecase) GetByUserID(userID string) (*model.Character, error) {
	return u.repo.GetByUserId(userID)
}

func (u *CharacterUsecase) LoginBonus(userID string) (*model.Character, error) {
	char, err := u.repo.GetByUserId(userID)
	if err != nil {
		return nil, err
	}
	service.CalcExp(char, 5)
	if err := u.repo.Update(char); err != nil {
		return nil, err
	}
	return char, nil
}
