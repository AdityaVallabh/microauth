package token

import (
	"bytes"
	"time"
)

type Cipher interface {
	Encrypt([]byte) []byte
	Decrypt([]byte) []byte
}

type SimpleTokenManager struct {
	Duration time.Duration
	Cipher   Cipher
}

type SimpleToken struct {
	Id          string
	ExpireAfter time.Time
}

// yes, I could use json encoder/decoder
func (t SimpleToken) Serialize() []byte {
	return []byte(t.Id + ":" + t.ExpireAfter.Format(time.RFC3339))
}

func (t *SimpleToken) Deserialize(b []byte) bool {
	email, expireAfter, ok := bytes.Cut(b, []byte(":"))
	if !ok {
		return false
	}
	exp, err := time.Parse(time.RFC3339, string(expireAfter))
	if err != nil {
		return false
	}
	t.Id = string(email)
	t.ExpireAfter = exp
	return true
}

func (t *SimpleTokenManager) GenerateToken(email string) (string, error) {
	token := SimpleToken{
		Id:          email,
		ExpireAfter: time.Now().Add(t.Duration),
	}
	return string(t.Cipher.Encrypt(token.Serialize())), nil
}

func (t *SimpleTokenManager) Validate(s string) (string, bool) {
	dec := t.Cipher.Decrypt([]byte(s))
	var token SimpleToken
	if !token.Deserialize(dec) {
		return "", false
	}
	if token.ExpireAfter.Before(time.Now()) {
		return "", false
	}
	return token.Id, true
}

func (t *SimpleTokenManager) Invalidate(s string) bool {
	return true
}

func (t *SimpleTokenManager) GetExpiry(s string) time.Time {
	dec := t.Cipher.Decrypt([]byte(s))
	var token SimpleToken
	if !token.Deserialize(dec) {
		return time.Time{}
	}
	return token.ExpireAfter
}
