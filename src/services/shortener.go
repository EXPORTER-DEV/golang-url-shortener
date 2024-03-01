package services

import (
	"golang-url-shortener/src/entity"
	"golang-url-shortener/src/repositories"
)

type ShortenerServiceInterface interface {
	Create(value string) (res *entity.ShortLink, err error)
	Get(key string) (res *entity.ShortLink, err error)
}

type ShortenerService struct {
	repository repositories.ShortenerRepositoryInterface
}

func (s *ShortenerService) Create(value string) (res *entity.ShortLink, err error) {
	return s.repository.AddShortLink(value)
}

func (s *ShortenerService) Get(key string) (res *entity.ShortLink, err error) {
	return s.repository.GetByShortLink(key)
}

func NewShortenerService(repository repositories.ShortenerRepositoryInterface) ShortenerServiceInterface {
	return &ShortenerService{repository}
}
