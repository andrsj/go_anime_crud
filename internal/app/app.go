package app

import (
	"github.com/labstack/echo/v4"

	"github.com/andrsj/go_anime_crud/internal/delivery/rest/api"
	"github.com/andrsj/go_anime_crud/internal/services/anime"
	"github.com/andrsj/go_anime_crud/pkg/logger"
)

type App struct {
	api           *api.API
	echo          *echo.Echo
	logger        logger.Interface
	anime_service anime.Interface
}

func (a *App) Run() error {
	a.logger.Info("Running server")

	err := a.echo.Start(":8080")
	if err != nil {
		a.logger.Fatal(err.Error())
	}

	return nil
}
