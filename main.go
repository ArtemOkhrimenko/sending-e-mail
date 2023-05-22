package main

import (
	"context"
	"email/intertal/api"
	"errors"
	"flag"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"os"

	"email/intertal/adapter"
	"email/intertal/app"
)

var (
	cfgFile = flag.String("cfg", "config.yml", "path to config file")
)

type config struct {
	Server server      `yaml:"server"`
	Email  emailConfig `yaml:"email"`
}
type server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
type emailConfig struct {
	APIKeyPublic  string `yaml:"api_key_public"`
	APIKeyPrivate string `yaml:"api_key_private"`
	Mailbox       string `yaml:"mailbox"`
}

func main() {
	ctx := context.Background()

	fmt.Println("Server is started..")

	err := start(ctx, *cfgFile)
	if err != nil {
		log.Fatal("shutdown", err)
	}
}

func start(ctx context.Context, configPath string) error {
	cfgFile, err := os.Open(configPath)
	if err != nil {
		return fmt.Errorf("os.Open: %w", err)
	}

	cfg := config{}
	err = yaml.NewDecoder(cfgFile).Decode(&cfg)
	if err != nil {
		return fmt.Errorf("yaml.NewDecoder.Decode: %w", err)
	}

	emailClient := mailing.New(mailing.Config{
		APIKeyPublic:   cfg.Email.APIKeyPublic,
		APIHostPrivate: cfg.Email.APIKeyPrivate,
		Mailbox:        cfg.Email.Mailbox,
	})

	application := app.New(emailClient)
	rout := api.New(application)

	srv := &http.Server{
		Addr:    cfg.Server.Port,
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

	return nil
}
