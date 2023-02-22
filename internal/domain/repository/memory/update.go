package memory

import (
	"fmt"

	"github.com/andrsj/go_anime_crud/internal/domain/model"
	"github.com/andrsj/go_anime_crud/internal/domain/repository"
)

func (s *RepositoryInMemory) UpdateAnimeCharacter(id model.IdAC, a *model.AnimeCharacter) (*model.AnimeCharacter, error) {
	err := validateNonZeroId(s.logger, id)
	if err != nil {
		return nil, err
	}

	_, ok := s.mapAnimaCharacter[id]
	if !ok {
		err := &repository.ACNotFoundError{Id: id}
		s.logger.Error(err.Error())
		return nil, err
	}

	s.logger.Info(fmt.Sprintf("Updating Anime Character by id: '%d'", id))
	s.mapAnimaCharacter[id] = a
	return a, nil
}
