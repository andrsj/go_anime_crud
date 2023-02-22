package memory

import (
	"fmt"

	"github.com/andrsj/go_anime_crud/internal/domain/model"
)

func (s *RepositoryInMemory) CreateAnimeCharacter(a *model.AnimeCharacter) model.IdAC {
	if a.Id != 0 {
		s.logger.Warn("Ignoring ID of Anime Character")
	}
	s.logger.Info(fmt.Sprintf("Adding Anime Character to memory with id: '%d'", s.IndexAC))
	id := s.IndexAC
	a.Id = id
	s.mapAnimaCharacter[id] = a

	s.IndexAC++

	s.logger.Info(fmt.Sprintf("Returning id of Anime Character in memory: '%d'", id))
	return id
}
