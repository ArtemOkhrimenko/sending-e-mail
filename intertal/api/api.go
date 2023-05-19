package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Application interface {
	SendEmail(ctx context.Context, toEmail string, title string, text string) error
}

type api struct {
	app Application
}

func New(app Application) http.Handler {
	router := gin.New()

	api := api{
		app: app,
	}

	numberHandler := router.Group("/contact")

	{
		numberHandler.GET("/", api.SendEmail)
	}

	return router
}
