package hasher

import (
	"crypto/sha256"
	"encoding/hex"
)

type Sha2 struct{}

func (h Sha2) Hash(s string) string {
	sha := sha256.New()
	sha.Write([]byte(s))
	return hex.EncodeToString(sha.Sum(nil))
}

func (h Sha2) CompareWithHash(s, hash string) bool {
	return h.Hash(s) == hash
}
