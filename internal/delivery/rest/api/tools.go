package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/andrsj/go_anime_crud/internal/domain/model"
)

func (e *API) Status(ctx echo.Context) error {
	e.logger.Info("Got a request :)")

	s := "Server says hello"

	e.logger.Info("Sending data to user")
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
		a.logger.Error(err.Error())
		return 0, err
	}
	return model.IdAC(id), nil
}

func handleError(ctx echo.Context, code int, err error) error {
	return ctx.JSON(code, map[string]string{"error": err.Error()})
}
