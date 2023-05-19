package app

import "context"

type App struct {
	email Email
}

type Email interface {
	SendEmail(ctx context.Context, toEmail, title, text string) error
}

func New() *App {
	return &App{}
}
