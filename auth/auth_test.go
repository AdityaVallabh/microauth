package auth

// import (
// 	"encoding/hex"
// 	"testing"
// 	"time"
// )

// type cipher struct{}

// func (c cipher) Encrypt(ptxt []byte) []byte {
// 	return []byte(hex.EncodeToString(ptxt))
// }

// func (c cipher) Decrypt(ctxt []byte) []byte {
// 	ptxt, _ := hex.DecodeString(string(ctxt))
// 	return ptxt
// }

// func TestAuther_GenerateTokenAndValidate(t *testing.T) {
// 	tests := []struct {
// 		email string
// 	}{
// 		{"hahaha@aswd.sd"},
// 	}

// 	a := NewAuther(nil, nil, new(cipher), time.Second)
// 	for _, tt := range tests {
// 		t.Run(tt.email, func(t *testing.T) {
// 			token, err := a.GenerateToken(tt.email)
// 			if err != nil {
// 				t.Errorf("Auther.GenerateToken() error = %v", err)
// 				return
// 			}
// 			email, ok := a.Validate(token)
// 			if !ok {
// 				t.Errorf("Auther.Validate() ok = %v", ok)
// 				return
// 			}
// 			if email != tt.email {
// 				t.Errorf("Auther.GenerateToken() = %v, want %v", email, tt.email)
// 			}
// 		})
// 	}
// }
