package repository

import "github.com/kakebon/backend/domain/model"

type UserRepository interface {
	Create(user *model.User) error
	GetByEmail (email string) (*model.User, error)
	Login(email string, password string) (*model.User, error)
}