package anime

import (
	"fmt"

	"github.com/andrsj/go_anime_crud/internal/domain/model"
)

func (s *Service) UpdateAnimeCharacter(id model.IdAC, a *model.AnimeCharacter) (*model.AnimeCharacter, error) {
	s.logger.Info(fmt.Sprintf("Updating one Anime Character by id '%d'", id))
	AC, err := s.repository.UpdateAnimeCharacter(id, a)
	if err != nil {
		return nil, err
	}
	return AC, nil
}
