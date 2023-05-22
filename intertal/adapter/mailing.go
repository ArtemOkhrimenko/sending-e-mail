package mailing

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/mailjet/mailjet-apiv3-go/v4"
)

const (
	timeout = 15 * time.Second
)

type (
	// Config for connect email client.
	Config struct {
		APIKeyPublic   string
		APIHostPrivate string
		Mailbox        string
	}

	// Client for send email.
	Client struct {
		cfg    Config
		client *mailjet.Client
	}
)

// New build and returns new client for send email.
func New(cfg Config) *Client {
	client := &http.Client{
		Timeout: timeout,
	}

	mailjetClient := mailjet.NewMailjetClient(cfg.APIKeyPublic, cfg.APIHostPrivate)
	mailjetClient.SetClient(client)

	return &Client{
		cfg:    cfg,
		client: mailjetClient,
	}
}

// SendEmail send email.
func (c *Client) SendEmail(ctx context.Context, toEmail, title, text string) error {
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: c.cfg.Mailbox,
				Name:  "Harvester",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: toEmail,
				},
			},
			Subject:  title,
			TextPart: title + "\n" + toEmail + "\n" + text,
		},
	}

	messages := mailjet.MessagesV31{Info: messagesInfo}

	_, err := c.client.SendMailV31(&messages, mailjet.WithContext(ctx))
	if err != nil {
		return fmt.Errorf("c.client.SendMailV31: %w", err)
	}

	return nil
}
