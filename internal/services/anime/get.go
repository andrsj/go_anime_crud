package anime

import (
	"fmt"

	"github.com/andrsj/go_anime_crud/internal/domain/model"
)

func (s *Service) GetAnimeCharacter(id model.IdAC) (*model.AnimeCharacter, error) {
	s.logger.Info(fmt.Sprintf("Getting one Anime Character by id '%d'", id))
	AC, err := s.repository.GetAnimeCharacter(id)
	if err != nil {
		return nil, err
	}
	return AC, nil
}

func (s *Service) GetAllAnimeCharacters() []*model.AnimeCharacter {
	s.logger.Info("Getting all Anime Character")
	sliceAC := s.repository.GetAllAnimeCharacters()
	return sliceAC
}
