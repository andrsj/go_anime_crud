package api

import (
	"net/http"

	"github.com/andrsj/go_anime_crud/internal/domain/model"
	"github.com/labstack/echo/v4"
)

// DELETE /api/ac/:id
func (a *API) DeleteAnimeCharacter(ctx echo.Context) error {
	id, err := a.parseID(ctx)
	if err != nil {
		return handleError(ctx, http.StatusBadRequest, err)
	}
	err = a.service.DeleteAnimeCharacter(model.IdAC(id))
	if err != nil {
		return handleError(ctx, http.StatusBadRequest, err)
	}
	return ctx.NoContent(http.StatusNoContent)
}
