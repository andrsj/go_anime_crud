package api

import (
	"github.com/andrsj/go_anime_crud/internal/services/anime"
	"github.com/andrsj/go_anime_crud/pkg/logger"
)

type API struct {
	service anime.Interface
	logger  logger.Interface
}

func New(s anime.Interface, l logger.Interface) *API {
	return &API{
		service: s,
		logger:  l,
	}
}
