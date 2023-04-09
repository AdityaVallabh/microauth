package token

import (
	"encoding/hex"
	"math/rand"
	"time"
)

type Store interface {
	Save(any, string) error
	Find(any, string, string) error
	Delete(any, string) error
}

type PersistedToken struct {
	Token  string `gorm:"primaryKey"`
	Id     string
	Expiry time.Time
}

type PersistedTokenManager struct {
	Rand     rand.Rand
	Store    Store
	Duration time.Duration
}

func (t *PersistedTokenManager) GenerateToken(id string) (string, error) {
	buf := make([]byte, 128)
	t.Rand.Read(buf)
	token := PersistedToken{
		Token:  hex.EncodeToString(buf),
		Id:     id,
		Expiry: time.Now().Add(t.Duration),
	}
	if err := t.Store.Save(&token, token.Token); err != nil {
		return "", err
	}
	return token.Token, nil
}

func (t *PersistedTokenManager) Validate(s string) (string, bool) {
	var token PersistedToken
	if err := t.Store.Find(&token, "token", s); err != nil {
		return "", false
	}
	return token.Id, true
}

func (t *PersistedTokenManager) Invalidate(s string) bool {
	return t.Store.Delete(&PersistedToken{}, s) == nil
}

func (t *PersistedTokenManager) GetExpiry(s string) time.Time {
	var token PersistedToken
	if err := t.Store.Find(&token, "token", s); err != nil {
		return time.Time{}
	}
	return token.Expiry
}
