package usecase

import "github.com/kakebon/backend/domain/repository"

type SlackUsecase struct {
	notifier repository.Notifier
}

func NewSlackUsecase(n repository.Notifier) *SlackUsecase {
	return &SlackUsecase{notifier: n}
}

func (u *SlackUsecase) Execute (message string) error {
	return u.notifier.Send(message)
}