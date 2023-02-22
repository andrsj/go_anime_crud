package app

import (
	"errors"

	"github.com/labstack/echo/v4"

	"github.com/andrsj/go_anime_crud/internal/delivery/rest/api"
	"github.com/andrsj/go_anime_crud/internal/services/anime"
	"github.com/andrsj/go_anime_crud/pkg/logger"
)

type AppBuilder struct {
	api           *api.API
	echo          *echo.Echo
	logger        logger.Interface
	anime_service anime.Interface
}

func NewAppBuilder() *AppBuilder {
	return &AppBuilder{}
}

func (b *AppBuilder) WithLogger(logger logger.Interface) *AppBuilder {
	b.logger = logger
	return b
}

func (b *AppBuilder) WithService(service anime.Interface) *AppBuilder {
	b.anime_service = service
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
	if b.logger == nil || b.anime_service == nil || b.api == nil || b.echo == nil {
		return nil, errors.New("logger, service, api and echo must be set to build App")
	}

	return &App{
		api:           b.api,
		echo:          b.echo,
		logger:        b.logger,
		anime_service: b.anime_service,
	}, nil
}
