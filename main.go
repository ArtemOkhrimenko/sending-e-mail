package main

import (
	"context"
	"email/intertal/api"
	"errors"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"log"
	"net/http"

	"email/intertal/adapter"
	"email/intertal/app"
)

func main() {
	ctx := context.Background()

	fmt.Println("Server is started..")

	//вынести в cfg?
	_ = mailing.New(mailing.Config{
		APIKeyPublic:   "",
		APIHostPrivate: "",
		Mailbox:        "",
	})

	application := app.New()
	rout := api.New(application)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: rout,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Fatal("listen: %w", err)
		}
	}()

	select {
	case <-ctx.Done():
		if err := srv.Shutdown(context.Background()); err != nil {
			panic(fmt.Sprintf("srv.Shutdown: %s", err.Error()))
		}
	}
}
