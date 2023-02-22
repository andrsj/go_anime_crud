package memory

import (
	"github.com/andrsj/go_anime_crud/internal/domain/model"
	"github.com/andrsj/go_anime_crud/internal/domain/repository"
	"github.com/andrsj/go_anime_crud/pkg/logger"
)

func validateNonZeroId(l logger.Interface, id model.IdAC) error {
	if id == 0 {
		err := &repository.ZeroIdError{}
		l.Error(err.Error())
		return err
	}
	return nil
}
