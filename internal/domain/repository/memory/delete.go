package memory

import (
	"fmt"

	"github.com/andrsj/go_anime_crud/internal/domain/model"
	"github.com/andrsj/go_anime_crud/internal/domain/repository"
)

func (s *RepositoryInMemory) DeleteAnimeCharacter(id model.IdAC) error {
	err := validateNonZeroId(s.logger, id)
	if err != nil {
		return err
	}

	_, ok := s.mapAnimaCharacter[id]
	if !ok {
		err := &repository.ACNotFoundError{Id: id}
		s.logger.Error(err.Error())
		return err
	}

	s.logger.Info(fmt.Sprintf("Deleting Anime Character by id: '%d'", id))
	delete(s.mapAnimaCharacter, id)

	return nil
}
