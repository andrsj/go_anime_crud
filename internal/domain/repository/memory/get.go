package memory

import (
	"fmt"

	"github.com/andrsj/go_anime_crud/internal/domain/model"
	"github.com/andrsj/go_anime_crud/internal/domain/repository"
)

func (s *RepositoryInMemory) GetAnimeCharacter(id model.IdAC) (*model.AnimeCharacter, error) {
	err := validateNonZeroId(s.logger, id)
	if err != nil {
		return nil, err
	}

	AC, ok := s.mapAnimaCharacter[id]
	if !ok {
		err := &repository.ACNotFoundError{Id: id}
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info(fmt.Sprintf("Returning Anime Character by id: '%d'", id))
	return AC, nil
}

func (s *RepositoryInMemory) GetAllAnimeCharacters() []*model.AnimeCharacter {
	result := make([]*model.AnimeCharacter, 0, len(s.mapAnimaCharacter))

	s.logger.Info("Iterating in all Anime Characters in memory")

	for key := range s.mapAnimaCharacter {
		result = append(result, s.mapAnimaCharacter[key])
	}

	s.logger.Info("Returning all Anime Characters in memory")
	return result
}
