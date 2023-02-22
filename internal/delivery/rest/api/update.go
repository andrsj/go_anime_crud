package api

import (
	"net/http"

	"github.com/andrsj/go_anime_crud/internal/domain/model"
	"github.com/labstack/echo/v4"
)

// PUT /api/ac/:id
func (a *API) UpdateAnimeCharacter(ctx echo.Context) error {
	id, err := a.parseID(ctx)
	if err != nil {
		return handleError(ctx, http.StatusBadRequest, err)
	}
	ac := new(model.AnimeCharacter)
	if err = ctx.Bind(ac); err != nil {
		return handleError(ctx, http.StatusBadRequest, err)
	}
	ac, err = a.service.UpdateAnimeCharacter(model.IdAC(id), ac)
	if err != nil {
		return handleError(ctx, http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, ac)
}
