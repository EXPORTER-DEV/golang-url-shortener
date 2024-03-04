package services

import (
	"fmt"
	"golang-url-shortener/common"
	"golang-url-shortener/entity"
	"golang-url-shortener/repositories"
)

type ShortenerServiceInterface interface {
	Create(value string) (res *entity.ShortLink, err error)
	Get(key string) (res *entity.ShortLink, err error)
}

type ShortenerService struct {
	repository repositories.ShortenerRepositoryInterface
}

func (s *ShortenerService) Create(value string) (res *entity.ShortLink, err error) {
	exists, err := s.repository.GetByShortLink(value)

	if err != nil {
		return nil, fmt.Errorf("failed while lookup existing shortlink in storage: %w", common.ErrDatabase)
	}

	if exists != nil {
		return nil, fmt.Errorf("short link already exists in database: %w", common.ErrDuplicate)
	}

	return s.repository.AddShortLink(value)
}

func (s *ShortenerService) Get(key string) (res *entity.ShortLink, err error) {
	return s.repository.GetByKey(key)
}

func NewShortenerService(repository repositories.ShortenerRepositoryInterface) ShortenerServiceInterface {
	return &ShortenerService{repository}
}
