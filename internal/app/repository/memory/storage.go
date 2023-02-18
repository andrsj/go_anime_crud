package memory

import (
	"fmt"

	"github.com/andrsj/go_anime_crud/internal/app/model"
	"github.com/andrsj/go_anime_crud/internal/app/repository"
	"github.com/andrsj/go_anime_crud/pkg/logger"
)

type RepositoryInMemory struct {
	logger  logger.Interface
	mapAC   repository.MapAnimeCharacters
	IndexAC model.IdAC
}

func New(l logger.Interface) (*RepositoryInMemory, error) {
	mapAC := make(repository.MapAnimeCharacters)
	s := &RepositoryInMemory{
		logger:  l,
		mapAC:   mapAC,
		IndexAC: 1,
	}
	return s, nil
}

func validateNonZeroId(l logger.Interface, id model.IdAC) error {
	if id == 0 {
		err := &repository.ZeroIdError{}
		l.Error(err.Error())
		return err
	}
	return nil
}

func (s *RepositoryInMemory) GetAnimeCharacter(id model.IdAC) (*model.AnimeCharacter, error) {
	err := validateNonZeroId(s.logger, id)
	if err != nil {
		return nil, err
	}

	AC, ok := s.mapAC[id]
	if !ok {
		err := &repository.ACNotFoundError{Id: id}
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info(fmt.Sprintf("Returning Anime Character by id: '%d'", id))
	return AC, nil
}

func (s *RepositoryInMemory) GetAllAnimeCharacters() []*model.AnimeCharacter {
	result := make([]*model.AnimeCharacter, 0, len(s.mapAC))

	s.logger.Info("Iterating in all Anime Characters in memory")

	for key := range s.mapAC {
		result = append(result, s.mapAC[key])
	}

	s.logger.Info("Returning all Anime Characters in memory")
	return result
}

func (s *RepositoryInMemory) CreateAnimeCharacter(a *model.AnimeCharacter) model.IdAC {
	if a.Id != 0 {
		s.logger.Warn("Ignoring ID of Anime Character")
	}
	s.logger.Info(fmt.Sprintf("Adding Anime Character to memory with id: '%d'", s.IndexAC))
	id := s.IndexAC
	a.Id = id
	s.mapAC[id] = a

	s.IndexAC++

	s.logger.Info(fmt.Sprintf("Returning id of Anime Character in memory: '%d'", id))
	return id
}

func (s *RepositoryInMemory) UpdateAnimeCharacter(id model.IdAC, a *model.AnimeCharacter) (*model.AnimeCharacter, error) {
	err := validateNonZeroId(s.logger, id)
	if err != nil {
		return nil, err
	}

	_, ok := s.mapAC[id]
	if !ok {
		err := &repository.ACNotFoundError{Id: id}
		s.logger.Error(err.Error())
		return nil, err
	}

	s.logger.Info(fmt.Sprintf("Updating Anime Character by id: '%d'", id))
	s.mapAC[id] = a
	return a, nil
}

func (s *RepositoryInMemory) DeleteAnimeCharacter(id model.IdAC) error {
	err := validateNonZeroId(s.logger, id)
	if err != nil {
		return err
	}

	_, ok := s.mapAC[id]
	if !ok {
		err := &repository.ACNotFoundError{Id: id}
		s.logger.Error(err.Error())
		return err
	}

	s.logger.Info(fmt.Sprintf("Deleting Anime Character by id: '%d'", id))
	delete(s.mapAC, id)

	return nil
}
