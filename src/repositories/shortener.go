package repositories

import (
	"golang-url-shortener/src/databases"
	"golang-url-shortener/src/entity"

	"github.com/go-redis/redis"
)

type ShortenerRepositoryInterface interface {
	GetByKey(key string) (*entity.ShortLink, error)
	GetByShortLink(value string) (*entity.ShortLink, error)
	AddShortLink(value string) (*entity.ShortLink, error)
}

type ShortenerRepository struct {
	database databases.RedisDatabaseInterface
}

func (r *ShortenerRepository) GetByKey(key string) (*entity.ShortLink, error) {
	data, err := r.database.Get(key)

	if err == redis.Nil {
		return nil, nil
	}

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

func (r *ShortenerRepository) GetByShortLink(value string) (*entity.ShortLink, error) {
	shortLink := entity.NewShortLink(value)

	data, err := r.database.Get(shortLink.ID)

	if err == redis.Nil {
		return nil, nil
	}

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

	return entity.NewShortLinkFromDatabase(value, shortLinkDatabase), nil
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
