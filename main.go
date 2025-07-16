package main

import (
	"fmt"

	"github.com/niteshghimire/cryptit/decrypt"
	"github.com/niteshghimire/cryptit/encrypt"
)

func main() {
	encryptedStr := encrypt.Nimbus("KodeKloud")
	fmt.Println(encryptedStr)
	fmt.Println(decrypt.Nimbus(encryptedStr))
}
