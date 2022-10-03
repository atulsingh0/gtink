package gtink

import (
	"log"

	"github.com/google/tink/go/aead"
	"github.com/google/tink/go/keyset"
)

func Generate() *keyset.Handle {
	kh, err := keyset.NewHandle(aead.AES128GCMKeyTemplate())
	if err != nil {
		log.Fatal(err)
	}
	return kh
}
