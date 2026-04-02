package usecase

import (
	"github.com/kakebon/backend/domain/model"
	"github.com/kakebon/backend/domain/repository"
	"github.com/kakebon/backend/domain/service"
)

type ExpenseUsecase struct {
	expenseRepo   repository.ExpenseRepository
	characterRepo repository.CharacterRepository
}

func NewExpenseUsecase(er repository.ExpenseRepository, cr repository.CharacterRepository) *ExpenseUsecase {
	return &ExpenseUsecase{expenseRepo: er, characterRepo: cr}
}

func (u *ExpenseUsecase) RecordExpense(expense *model.Expense) (*model.Character, error) {
	if err := u.expenseRepo.Create(expense); err != nil {
		return nil, err
	}
	char, err := u.characterRepo.GetByUserId(expense.UserID)
	if err != nil {
		return nil, err
	}
	service.CalcExp(char, expense.Amount/100)
	if err := u.characterRepo.Update(char); err != nil {
		return nil, err
	}
	return char, nil
}

func (u *ExpenseUsecase) GetAll() ([]model.Expense, error) {
	return u.expenseRepo.GetAll()
}
