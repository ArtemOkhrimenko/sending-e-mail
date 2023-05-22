package app

import "context"

type App struct {
	email Email
}

type Email interface {
	SendEmail(ctx context.Context, text string) error
}

func New(email Email) *App {
	return &App{
		email: email,
	}
}
