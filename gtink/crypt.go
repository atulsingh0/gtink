package gtink

import (
	"github.com/google/tink/go/aead"
	"github.com/google/tink/go/keyset"
	"github.com/google/tink/go/tink"
)


func Encrypt(kh *keyset.Handle, msg string, hint string) ([]byte, tink.AEAD) {

	a,_:= aead.New(kh)
	ct, _ := a.Encrypt([]byte(msg), []byte(hint))

	return ct, a
}

func Decrypt(a tink.AEAD, encryptText string, hint string) ([]byte) {

	pt, _ := a.Decrypt([]byte(encryptText), []byte(hint))

	return pt
}