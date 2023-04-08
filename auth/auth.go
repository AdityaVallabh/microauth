package auth

import "time"

type Storage interface {
	PasswordHashByEmail(email string) (phash string, err error)
	SetPasswordHashByEmail(email string, phash string) error
}

type Hasher interface {
	Hash(string) string
	CompareWithHash(s string, hash string) bool
}

type TokenManager interface {
	GenerateToken(email string) (string, error)
	Validate(token string) (string, bool)
	Invalidate(token string) bool
	GetExpiry(token string) time.Time
}

type Auther struct {
	TokenManager
	Hasher  Hasher
	Storage Storage
}

func (a Auther) CheckPassword(email, password string) bool {
	h, err := a.Storage.PasswordHashByEmail(email)
	if err != nil {
		return false
	}
	return a.Hasher.CompareWithHash(password, h)
}

func (a Auther) SavePassword(email, password string) error {
	hash := a.Hasher.Hash(password)
	return a.Storage.SetPasswordHashByEmail(email, hash)
}
