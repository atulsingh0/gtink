package main

import (
	"fmt"

	"github.com/atulsingh0/gtink/gtink"
)

const (
	msg  string = "Exploring gtink"
	hint string = "this is a test"
)

func main() {
	fmt.Println("Hello World !")

	keyset := gtink.Generate()
	fmt.Println(keyset)

	fmt.Printf("Message: %s\nAssociated data: %s\n", msg, hint)

	ct := gtink.Encrypt(keyset, msg, hint)
	pt := gtink.Decrypt(keyset, string(ct), hint)

	fmt.Printf("Cipher text: %x\nPlain text: %s\n\n\n", ct, pt)

}
