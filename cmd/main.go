package main

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/andrsj/go_anime_crud/internal/app"
	"github.com/andrsj/go_anime_crud/internal/repository/memory"
	"github.com/andrsj/go_anime_crud/internal/rest/api"
	"github.com/andrsj/go_anime_crud/internal/service"
	"github.com/andrsj/go_anime_crud/pkg/logger/zerolog"
)

func main() {

	// Setup logger
	l := zerolog.New()

	// Setup service
	repo, err := memory.New(l)
	if err != nil {
		return
	}

	s, err := service.New(l, repo)
	if err != nil {
		l.Error("Error while setting up the services")
		return
	}

	// Setup server side handler
	a := api.New(s, l)

	l.Info("Setting up the Echo instance")
	e := echo.New()

	e.GET("/", a.Status)

	e.POST("/api/ac", a.CreateAnimeCharacter)
	e.GET("/api/ac/:id", a.GetAnimeCharacter)
	e.GET("/api/ac/", a.GetAllAnimeCharacters)
	e.PUT("/api/ac/:id", a.UpdateAnimeCharacter)
	e.DELETE("/api/ac/:id", a.DeleteAnimeCharacter)

	app, err := app.NewAppBuilder().
		WithLogger(l).
		WithService(s).
		WithAPI(a).
		WithEcho(e).
		Build()

	if err != nil {
		l.Fatal(fmt.Sprintf("App can't be created, error: %s", err))
	}

	if err := app.Run(); err != nil {
		l.Fatal(fmt.Sprintf("App exited with error: %s", err))
	}
}
