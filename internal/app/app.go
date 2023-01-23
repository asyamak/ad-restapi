package app

import (
	"ad-api/config"
	"ad-api/internal/repository"
	database "ad-api/pkg/database"
	"fmt"
)


type App struct{
	config *config.Config

}

func New(config *config.Config) *App{
	return &App{
		 config,
	}

}

func(a *App) Start() error{

	db, err := database.New(a.config)
	if err != nil {
		return fmt.Errorf("app: start: db: %w",err)
	}
	
	AdsRepository := repository.NewAdRepository(db)
	fmt.Println(AdsRepository)

	return nil



}