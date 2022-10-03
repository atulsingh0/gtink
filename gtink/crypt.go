package gtink

import (
	"github.com/google/tink/go/aead"
	"github.com/google/tink/go/keyset"
)

func Encrypt(kh *keyset.Handle, msg string, hint string) []byte {

	a, _ := aead.New(kh)
	ct, _ := a.Encrypt([]byte(msg), []byte(hint))

	return ct
}

func Decrypt(kh *keyset.Handle, encryptText string, hint string) []byte {

	a, _ := aead.New(kh)
	pt, _ := a.Decrypt([]byte(encryptText), []byte(hint))

	return pt
}
