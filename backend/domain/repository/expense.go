package repository

import "github.com/kakebon/backend/domain/model"

type ExpenseRepository interface {
	Create(expense *model.Expense) error
	GetAll() ([]model.Expense, error)
	Delete(id string) error
	GetByID(id string) (*model.Expense, error)
	Update(id string, expese *model.Expense) (*model.Expense, error)		
}