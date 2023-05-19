package app

import (
	"context"
	"fmt"
)

func (a *App) SendEmail(ctx context.Context, toEmail string, title string, text string) error {
	err := a.email.SendEmail(ctx, toEmail, title, text)
	if err != nil {
		return fmt.Errorf("a.email.SendEmail: %w", err)
	}

	return nil
}
