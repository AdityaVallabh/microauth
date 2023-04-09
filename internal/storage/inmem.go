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

func (s *Storage) Find(v any, _, keyValue string) error {
	entry, err := s.Get(storage.Key(keyValue))
	if err != nil {
		return err
	}
	b, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}

func (s *Storage) Save(v any, keyValue string) error {
	var entry storage.Entry
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &entry)
	if err != nil {
		return err
	}
	return s.Set(storage.Key(keyValue), entry)
}

func NewInMem() *Storage {
	return &Storage{storage.NewInMem()}
}
