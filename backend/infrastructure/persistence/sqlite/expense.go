package sqlite

import (
	"fmt"
	"net/http"

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

func (p *expensePersistence) Delete(id string) error {
	res := p.DB.Where("id = ?", id).Delete(&model.Expense{}); 
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("record not found")
	}
	return nil
}

func (p *expensePersistence) GetByID(id string) (*model.Expense, error) {
	var res model.Expense
	
	if err := p.DB.Where("id = ?", id).First(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}