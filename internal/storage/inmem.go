package storage

import (
	"encoding/json"
	"errors"
	"microauth/pkg/storage"
	"reflect"
)

type Storage struct {
	*storage.InMem
}

func (s *Storage) AutoMigrate(v ...any) error {
	for _, t := range v {
		s.Set(s.getType(t), storage.Entry{})
	}
	return nil
}

func (s *Storage) Find(v any, _, keyValue string) error {
	entry, err := s.Get(s.getType(v))
	if err != nil {
		return err
	}
	val, ok := entry[keyValue]
	if !ok {
		return errors.New("not found")
	}
	b, err := json.Marshal(val)
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
	db, _ := s.Get(s.getType(v))
	db[keyValue] = entry
	return nil
}

func (s *Storage) getType(v any) storage.Key {
	return storage.Key(reflect.ValueOf(v).Type().String())
}

func NewInMem() *Storage {
	return &Storage{storage.NewInMem()}
}
