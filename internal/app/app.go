package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"ad-api/config"
	"ad-api/internal/controllers"
	"ad-api/internal/repository"
	"ad-api/internal/server"
	"ad-api/internal/usecase"
	database "ad-api/pkg/database"
)

type App struct {
	config *config.Config
}

func New(config *config.Config) *App {
	return &App{
		config,
	}
}

func (a *App) Start() error {
	db, err := database.New(a.config)
	if err != nil {
		return fmt.Errorf("app: start: db: %w", err)
	}

	adsRepository := repository.NewAdRepository(db)
	adsUsecase := usecase.NewUsecase(adsRepository)
	adsHandler := controllers.NewHandler(adsUsecase)

	router := controllers.SetUpRouter(adsHandler)
	server := server.New(a.config, router, db)

	// check graceful shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Printf("signal: " + s.String())
	case err = <-server.Notify():
		log.Printf("signal.Notify: %v", err)
	}

	err = server.Shutdown()
	if err != nil {
		log.Printf("server shutdown: %v", err)
		return err
	}
	return nil
}
