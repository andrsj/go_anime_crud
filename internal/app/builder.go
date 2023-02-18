package app

import (
	"errors"

	"github.com/andrsj/go_anime_crud/internal/app/rest/api"
	"github.com/andrsj/go_anime_crud/internal/app/service"
	"github.com/andrsj/go_anime_crud/pkg/logger"
	"github.com/labstack/echo/v4"
)

type AppBuilder struct {
	logger  logger.Interface
	service service.Interface
	api     *api.API
	echo    *echo.Echo
}

func NewAppBuilder() *AppBuilder {
	return &AppBuilder{}
}

func (b *AppBuilder) WithLogger(logger logger.Interface) *AppBuilder {
	b.logger = logger
	return b
}

func (b *AppBuilder) WithService(service service.Interface) *AppBuilder {
	b.service = service
	return b
}

func (b *AppBuilder) WithAPI(api *api.API) *AppBuilder {
	b.api = api
	return b
}

func (b *AppBuilder) WithEcho(echo *echo.Echo) *AppBuilder {
	b.echo = echo
	return b
}

func (b *AppBuilder) Build() (*App, error) {
	if b.logger == nil || b.service == nil || b.api == nil || b.echo == nil {
		return nil, errors.New("logger, service, api and echo must be set to build App")
	}

	return &App{
		logger:  b.logger,
		service: b.service,
		api:     b.api,
		echo:    b.echo,
	}, nil
}
