package cipher

import "encoding/hex"

type HexCipher struct{}

func (c HexCipher) Encrypt(ptxt []byte) []byte {
	return []byte(hex.EncodeToString(ptxt))
}

func (c HexCipher) Decrypt(ctxt []byte) []byte {
	ptxt, _ := hex.DecodeString(string(ctxt))
	return ptxt
}
