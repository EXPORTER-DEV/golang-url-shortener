package repositories

import (
	"golang-url-shortener/src/databases"
	"golang-url-shortener/src/entity"
	"testing"
)

type MockedRedisDatabase struct {
	storage map[string]string
}

func (d *MockedRedisDatabase) Get(key string) (value string, err error) {
	value, ok := d.storage[key]

	if !ok || value == "" {
		return "", nil
	}

	return value, nil
}

func (d *MockedRedisDatabase) Set(key string, value string) error {
	d.storage[key] = value

	return nil
}

func NewMockedRedisDatabase(storage map[string]string) databases.RedisDatabaseInterface {
	return &MockedRedisDatabase{
		storage,
	}
}

type ShortenerRepositoryCreationTestData struct {
	Value    string
	Expected *entity.ShortLink
}

func TestShortenerRepositoryCreation(t *testing.T) {
	value := "123456"

	testData := &ShortenerRepositoryCreationTestData{
		Value: value,
		Expected: &entity.ShortLink{
			ID:       string([]rune("8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92")),
			Value:    &value,
			Database: &entity.ShortLinkDatabase{},
		},
	}
	mockedRedisDatabase := NewMockedRedisDatabase(make(map[string]string))

	repository := NewShortenerRepository(mockedRedisDatabase)

	res, err := repository.AddShortLink(testData.Value)

	if err != nil {
		t.Log("Got unexpected error", err)
		t.Fail()
		return
	}

	if res.ID != testData.Expected.ID || *res.Value != *testData.Expected.Value {
		t.Logf("%x : %x", res.ID, testData.Expected.ID)
		t.Log("Got actual with expected mismatch", "actual:", res, "expected:", testData.Expected)
		t.Fail()
		return
	}

}
