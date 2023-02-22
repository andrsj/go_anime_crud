package api

import (
	"net/http"

	"github.com/andrsj/go_anime_crud/internal/domain/model"
	"github.com/labstack/echo/v4"
)

// POST /api/ac/
func (a *API) CreateAnimeCharacter(ctx echo.Context) error {
	ac := new(model.AnimeCharacter)
	if err := ctx.Bind(ac); err != nil {
		return handleError(ctx, http.StatusBadRequest, err)
	}
	a.service.CreateAnimeCharacter(ac)
	return ctx.JSON(http.StatusCreated, ac)
}
