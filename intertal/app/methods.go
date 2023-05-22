package app

import (
	"context"
	"errors"
	"fmt"
	"regexp"
)

func (a *App) SendEmail(ctx context.Context, name, contact, text string) error {
	disc, err := validate(name, contact, text)
	if err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	err = a.email.SendEmail(ctx, disc)
	if err != nil {
		return fmt.Errorf("a.email.SendEmail: %w", err)
	}

	return nil
}

func validate(name, contact, disc string) (text string, err error) {
	emailRegex := regexp.MustCompile("^[a-z0-9._%+\\-]+@[a-z0-9.\\-]+\\.[a-z]{2,4}$")
	phoneRegex := regexp.MustCompile("^((8|\\+7)[\\- ]?)?(\\(?\\d{3}\\)?[\\- ]?)?[\\d\\- ]{7,10}$")
	telegramRegex := regexp.MustCompile("^[A-Za-z\\d_]{5,32}$")

	//text = "name: " + name + "\n" + "telegram: " + contact + "\n" + disc

	if telegramRegex.MatchString(contact) {
		text = "name: " + name + "\n" + "telegram: " + contact + "\n" + disc
		return text, nil
	}

	if emailRegex.MatchString(contact) {
		text = "name: " + name + "\n" + "email: " + contact + "\n" + disc
		return text, nil
	}

	if phoneRegex.MatchString(contact) {
		text = "name: " + name + "\n" + "phone: " + contact + "\n" + disc
		return text, nil
	}

	return "", errors.New("ErrorValidate")
}
