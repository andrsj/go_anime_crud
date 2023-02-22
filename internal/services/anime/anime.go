package anime

import (
	"github.com/andrsj/go_anime_crud/internal/domain/repository"
	"github.com/andrsj/go_anime_crud/pkg/logger"
)

type Interface interface {
	repository.Interface
}

type Service struct {
	logger     logger.Interface
	repository repository.Interface
}

func New(l logger.Interface, r repository.Interface) *Service {
	return &Service{
		logger:     l,
		repository: r,
	}
}
