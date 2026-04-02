package sqlite

import (
	"github.com/kakebon/backend/domain/model"
	"gorm.io/gorm"
)

type expensePersistence struct {
	DB *gorm.DB
}

func NewExpensePersistence(db *gorm.DB) *expensePersistence {
	return &expensePersistence{DB: db}
}

func (p *expensePersistence) Create(expense *model.Expense) error {
	return p.DB.Create(expense).Error
}

func (p *expensePersistence) GetAll() ([]model.Expense, error) {
	var expenses []model.Expense
	if err := p.DB.Find(&expenses).Error; err != nil {
		return nil, err
	}
	return expenses, nil
}