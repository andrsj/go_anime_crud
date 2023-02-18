package service

import (
	"fmt"

	"github.com/andrsj/go_anime_crud/internal/app/model"
	"github.com/andrsj/go_anime_crud/internal/app/repository"
	"github.com/andrsj/go_anime_crud/pkg/logger"
)

type Interface interface {
	repository.Interface
	Ping() string
}

type Service struct {
	logger     logger.Interface
	repository repository.Interface
}

func New(l logger.Interface, r repository.Interface) (*Service, error) {
	return &Service{
			logger:     l,
			repository: r,
		},
		nil
}

func (s *Service) Ping() string {
	s.logger.Info("Pinging...")
	return "Ping"
}

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

func (s *Service) CreateAnimeCharacter(a *model.AnimeCharacter) model.IdAC {
	s.logger.Info("Creating one Anime Character")
	id := s.repository.CreateAnimeCharacter(a)
	return id
}

func (s *Service) UpdateAnimeCharacter(id model.IdAC, a *model.AnimeCharacter) (*model.AnimeCharacter, error) {
	s.logger.Info(fmt.Sprintf("Updating one Anime Character by id '%d'", id))
	AC, err := s.repository.UpdateAnimeCharacter(id, a)
	if err != nil {
		return nil, err
	}
	return AC, nil
}

func (s *Service) DeleteAnimeCharacter(id model.IdAC) error {
	s.logger.Info(fmt.Sprintf("Deleting one Anime Character by id '%d'", id))
	err := s.repository.DeleteAnimeCharacter(id)
	if err != nil {
		return err
	}
	return nil
}
