package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GET /api/ac/:id
func (a *API) GetAnimeCharacter(ctx echo.Context) error {
	id, err := a.parseID(ctx)
	if err != nil {
		return handleError(ctx, http.StatusBadRequest, err)
	}
	ac, err := a.service.GetAnimeCharacter(id)
	if err != nil {
		return handleError(ctx, http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, ac)
}

// GET /api/ac/
func (a *API) GetAllAnimeCharacters(ctx echo.Context) error {
	slcAC := a.service.GetAllAnimeCharacters()
	return ctx.JSON(http.StatusOK, slcAC)
}
