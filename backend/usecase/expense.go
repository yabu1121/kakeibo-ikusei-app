package usecase

import (
	"log"
	"strconv"

	"github.com/kakebon/backend/domain/model"
	"github.com/kakebon/backend/domain/repository"
	"github.com/kakebon/backend/domain/service"
)

type ExpenseUsecase struct {
	expenseRepo   repository.ExpenseRepository
	characterRepo repository.CharacterRepository
	notifier repository.Notifier
}

func NewExpenseUsecase(er repository.ExpenseRepository, cr repository.CharacterRepository, n repository.Notifier) *ExpenseUsecase {
	return &ExpenseUsecase{expenseRepo: er, characterRepo: cr, notifier: n}
}

func (u *ExpenseUsecase) RecordExpense(expense *model.Expense) (*model.Character, error) {
	if err := u.expenseRepo.Create(expense); err != nil {
		return nil, err
	}
	char, err := u.characterRepo.GetByUserId(expense.UserID)
	if err != nil {
		return nil, err
	}

	if err := u.notifier.Send(strconv.Itoa(expense.Amount) + "使用しました"); err != nil {
		log.Printf("failed to send notification: %v", err)
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

func (u *ExpenseUsecase) Delete(id string) error {
	expense, err := u.expenseRepo.GetByID(id)
	if err != nil {
		return err
	}
	
	char, err := u.characterRepo.GetByUserId(expense.UserID)
	if err != nil {
		return err
	}

	exp := expense.Amount / 100
	if char.CurrentExp >= exp {
		char.CurrentExp -= exp
	} else {
		char.CurrentExp = 0
	}

	if err := u.characterRepo.Update(char); err != nil {
		return err
	}

	if err := u.expenseRepo.Delete(id); err != nil {
		return err
	}

	return nil
}

func (u *ExpenseUsecase) GetByID (id string) (*model.Expense, error) {
	return u.expenseRepo.GetByID(id)
}