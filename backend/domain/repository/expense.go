package repository

import "github.com/kakebon/backend/domain/model"

type ExpenseRepository interface {
	Create(expense *model.Expense) error
	GetAll() ([]model.Expense, error)
}