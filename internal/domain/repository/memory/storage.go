package memory

import (
	"github.com/andrsj/go_anime_crud/internal/domain/model"
	"github.com/andrsj/go_anime_crud/internal/domain/repository"
	"github.com/andrsj/go_anime_crud/pkg/logger"
)

type RepositoryInMemory struct {
	logger            logger.Interface
	IndexAC           model.IdAC
	mapAnimaCharacter repository.MapAnimeCharacters
}

func New(l logger.Interface) (*RepositoryInMemory, error) {
	mapAC := make(repository.MapAnimeCharacters)
	s := &RepositoryInMemory{
		logger:            l,
		IndexAC:           1,
		mapAnimaCharacter: mapAC,
	}
	return s, nil
}
