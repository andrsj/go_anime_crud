package repository

import (
	"fmt"

	"github.com/andrsj/go_anime_crud/internal/model"
)

type ACNotFoundError struct {
	Id model.IdAC
}

func (a *ACNotFoundError) Error() string {
	return fmt.Sprintf("not found AnimeCharacter by id '%d'", a.Id)
}

type ZeroIdError struct{}

func (z *ZeroIdError) Error() string {
	return "Id can't be zero (0)"
}
