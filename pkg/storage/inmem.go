package storage

import "errors"

type Key string

type Entry map[string]any

type InMem struct {
	db map[Key]Entry
}

func (m *InMem) Get(k Key) (Entry, error) {
	entry, ok := m.db[k]
	if !ok {
		return nil, errors.New("no such key")
	}
	return entry, nil
}

func (m *InMem) Set(k Key, v Entry) error {
	m.db[k] = v
	return nil
}

func NewInMem() *InMem {
	return &InMem{
		db: make(map[Key]Entry),
	}
}
