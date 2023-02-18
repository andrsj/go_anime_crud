package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/andrsj/go_anime_crud/internal/app/model"
	"github.com/andrsj/go_anime_crud/internal/app/service"
	"github.com/andrsj/go_anime_crud/pkg/logger"
	"github.com/labstack/echo/v4"
)

type API struct {
	s service.Interface
	l logger.Interface
}

func New(s service.Interface, l logger.Interface) *API {
	return &API{
		s: s,
		l: l,
	}
}

func (e *API) Status(ctx echo.Context) error {
	e.l.Info("Got a request :)")

	data := e.s.Ping()
	s := fmt.Sprintf("Server data is: %s", data)

	e.l.Info("Sending data to user")
	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return err
	}

	return nil
}

func (a *API) parseID(ctx echo.Context) (model.IdAC, error) {
	strID := ctx.Param("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		err = fmt.Errorf("ID can't be not 'int' value: [GO: %w]", err)
		a.l.Error(err.Error())
		return 0, err
	}
	return model.IdAC(id), nil
}

func handleError(ctx echo.Context, code int, err error) error {
	return ctx.JSON(code, map[string]string{"error": err.Error()})
}

// POST /api/ac/
func (a *API) CreateAnimeCharacter(ctx echo.Context) error {
	ac := new(model.AnimeCharacter)
	if err := ctx.Bind(ac); err != nil {
		return handleError(ctx, http.StatusBadRequest, err)
	}
	a.s.CreateAnimeCharacter(ac)
	return ctx.JSON(http.StatusCreated, ac)
}

// GET /api/ac/:id
func (a *API) GetAnimeCharacter(ctx echo.Context) error {
	id, err := a.parseID(ctx)
	if err != nil {
		return handleError(ctx, http.StatusBadRequest, err)
	}
	ac, err := a.s.GetAnimeCharacter(id)
	if err != nil {
		return handleError(ctx, http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, ac)
}

// GET /api/ac/
func (a *API) GetAllAnimeCharacters(ctx echo.Context) error {
	slcAC := a.s.GetAllAnimeCharacters()
	return ctx.JSON(http.StatusOK, slcAC)
}

// PUT /api/ac/:id
func (a *API) UpdateAnimeCharacter(ctx echo.Context) error {
	id, err := a.parseID(ctx)
	if err != nil {
		return handleError(ctx, http.StatusBadRequest, err)
	}
	ac := new(model.AnimeCharacter)
	if err := ctx.Bind(ac); err != nil {
		return handleError(ctx, http.StatusBadRequest, err)
	}
	ac, err = a.s.UpdateAnimeCharacter(model.IdAC(id), ac)
	if err != nil {
		return handleError(ctx, http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, ac)
}

// DELETE /api/ac/:id
func (a *API) DeleteAnimeCharacter(ctx echo.Context) error {
	id, err := a.parseID(ctx)
	if err != nil {
		return handleError(ctx, http.StatusBadRequest, err)
	}
	err = a.s.DeleteAnimeCharacter(model.IdAC(id))
	if err != nil {
		return handleError(ctx, http.StatusBadRequest, err)
	}
	return ctx.NoContent(http.StatusNoContent)
}
