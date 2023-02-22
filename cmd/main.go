package main

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/andrsj/go_anime_crud/internal/app"
	"github.com/andrsj/go_anime_crud/internal/delivery/rest/api"
	"github.com/andrsj/go_anime_crud/internal/domain/repository/memory"
	anime_service "github.com/andrsj/go_anime_crud/internal/services/anime"
	"github.com/andrsj/go_anime_crud/pkg/logger/zerolog"
)

func main() {

	log := zerolog.New()

	repository, err := memory.New(log)
	if err != nil {
		return
	}

	service_app := anime_service.New(log, repository)

	api_router := api.New(service_app, log)

	echo_router := echo.New()

	echo_router.GET("/", api_router.Status)

	echo_router.POST("/api/ac", api_router.CreateAnimeCharacter)
	echo_router.GET("/api/ac/:id", api_router.GetAnimeCharacter)
	echo_router.GET("/api/ac/", api_router.GetAllAnimeCharacters)
	echo_router.PUT("/api/ac/:id", api_router.UpdateAnimeCharacter)
	echo_router.DELETE("/api/ac/:id", api_router.DeleteAnimeCharacter)

	app, err := app.NewAppBuilder().
		WithLogger(log).
		WithService(service_app).
		WithAPI(api_router).
		WithEcho(echo_router).
		Build()

	if err != nil {
		log.Fatal(fmt.Sprintf("App can't be created, error: %s", err))
	}

	if err := app.Run(); err != nil {
		log.Fatal(fmt.Sprintf("App exited with error: %s", err))
	}
}
