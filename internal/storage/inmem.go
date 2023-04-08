package storage

import (
	"encoding/json"
	"microauth/pkg/storage"
)

type Storage struct {
	*storage.InMem
}

func (s *Storage) AutoMigrate(v ...any) error {
	return nil
}

func (s *Storage) Find(key string, v any) error {
	entry, err := s.Get(storage.Key(key))
	if err != nil {
		return err
	}
	b, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}

func (s *Storage) Save(v any) error {
	var entry storage.Entry
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &entry)
	if err != nil {
		return err
	}
	return s.Set(entry["email"].(storage.Key), entry)
}

func (s *Storage) PasswordHashByEmail(email string) (string, error) {
	entry, err := s.Get(storage.Key(email))
	if err != nil {
		return "", err
	}
	return entry["phash"].(string), nil
}

func (s *Storage) SetPasswordHashByEmail(email, phash string) error {
	entry := make(storage.Entry)
	entry["phash"] = phash
	return s.Set(storage.Key(email), entry)
}

func NewInMem() *Storage {
	return &Storage{storage.NewInMem()}
}
