package anime

import "github.com/andrsj/go_anime_crud/internal/domain/model"

func (s *Service) CreateAnimeCharacter(a *model.AnimeCharacter) model.IdAC {
	s.logger.Info("Creating one Anime Character")
	id := s.repository.CreateAnimeCharacter(a)
	return id
}
