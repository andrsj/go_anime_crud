package anime

import (
	"fmt"

	"github.com/andrsj/go_anime_crud/internal/domain/model"
)

func (s *Service) DeleteAnimeCharacter(id model.IdAC) error {
	s.logger.Info(fmt.Sprintf("Deleting one Anime Character by id '%d'", id))
	err := s.repository.DeleteAnimeCharacter(id)
	if err != nil {
		return err
	}
	return nil
}
