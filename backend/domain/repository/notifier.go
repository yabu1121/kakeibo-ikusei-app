package repository

type Notifier interface {
	Send (message string) error
}