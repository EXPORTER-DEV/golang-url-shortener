package entity

import (
	"crypto"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

const SHORTLINK_ID_LENGTH = 8

type ShortLinkDatabase struct {
	Value   string `json:"value"`
	Created int    `json:"created"`
}

func (d *ShortLinkDatabase) ToJSONString() (string, error) {
	res, err := json.Marshal(d)

	return string(res), err
}

type ShortLink struct {
	Database *ShortLinkDatabase
	Value    *string
	ID       string
}

func (d *ShortLink) String() string {
	return fmt.Sprintf("ShortLink{Value:%v, ID:%x}", *d.Value, d.ID)
}

func NewShortLinkDatabase(data string) (*ShortLinkDatabase, error) {
	shortLinkDatabase := &ShortLinkDatabase{}

	err := json.Unmarshal([]byte(data), shortLinkDatabase)

	return shortLinkDatabase, err
}

func NewShortLinkFromDatabase(ID string, shortLinkDatabase *ShortLinkDatabase) *ShortLink {
	shortLink := &ShortLink{
		Database: shortLinkDatabase,
		Value:    &shortLinkDatabase.Value,
		ID:       ID,
	}

	return shortLink
}

func NewShortLink(value string) *ShortLink {
	h := crypto.SHA256.New()
	h.Write([]byte(value))

	ID := hex.EncodeToString(h.Sum(nil))

	shortLinkDatabase := &ShortLinkDatabase{
		value,
		time.Now().Second(),
	}

	return &ShortLink{
		Database: shortLinkDatabase,
		Value:    &value,
		ID:       ID,
	}
}
