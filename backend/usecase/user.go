package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/kakebon/backend/domain/model"
	"github.com/kakebon/backend/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase (repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) Create (name, email, password string) (*model.User ,error) {
	newUserID := uuid.NewString()
	newCharID := uuid.NewString()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID: newUserID,
		Name: name,
		Email: email,
		HashedPassword: string(hashedPassword),
		Character: &model.Character{
			ID: newCharID,
			UserID: newUserID,
			CurrentLevel: 1,
			CurrentExp: 0,
			ImageURL: model.GetImageByLevel(1),
		},
	}

	if err := u.repo.Create(user); err != nil {
		return nil, fmt.Errorf("failed to create user with character")
	}

	return user, nil
}

func (u *UserUsecase) GetByEmail (email string) (*model.User, error) {
	user, err := u.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUsecase) Login(email string, password string) (*model.User, error) {
	user, err := u.repo.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("ログイン情報が違います")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword),[]byte(password)); err != nil {
		return nil, fmt.Errorf("ログイン情報が違います")
	}
	return user, nil
}