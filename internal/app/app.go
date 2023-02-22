package app

import (
	"github.com/andrsj/go_anime_crud/internal/rest/api"
	"github.com/andrsj/go_anime_crud/internal/service"
	"github.com/andrsj/go_anime_crud/pkg/logger"
	"github.com/labstack/echo/v4"
)

type App struct {
	logger  logger.Interface
	service service.Interface
	api     *api.API
	echo    *echo.Echo
}

func (a *App) Run() error {
	a.logger.Info("Running server")

	err := a.echo.Start(":8080")
	if err != nil {
		a.logger.Fatal(err.Error())
	}

	return nil
}
