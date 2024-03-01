package repositories

import (
	"golang-url-shortener/src/databases"
	"golang-url-shortener/src/entity"
)

type ShortenerRepositoryInterface interface {
	GetByShortLink(key string) (*entity.ShortLink, error)
	AddShortLink(value string) (*entity.ShortLink, error)
}

type ShortenerRepository struct {
	database databases.RedisDatabaseInterface
}

func (r *ShortenerRepository) GetByShortLink(key string) (*entity.ShortLink, error) {
	data, err := r.database.Get(key)

	if err != nil {
		return nil, err
	}

	if data == "" {
		return nil, nil
	}

	shortLinkDatabase, err := entity.NewShortLinkDatabase(data)

	if err != nil {
		return nil, err
	}

	return entity.NewShortLinkFromDatabase(key, shortLinkDatabase), nil
}

func (r *ShortenerRepository) AddShortLink(value string) (*entity.ShortLink, error) {
	shortLink := entity.NewShortLink(value)

	stringifiedJSON, err := shortLink.Database.ToJSONString()

	if err != nil {
		return nil, err
	}

	err = r.database.Set(shortLink.ID, stringifiedJSON)

	return shortLink, err
}

func NewShortenerRepository(database databases.RedisDatabaseInterface) ShortenerRepositoryInterface {
	return &ShortenerRepository{
		database,
	}
}
