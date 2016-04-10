package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
)

/**
 * Cryptopals rule
 * Always operate on raw bytes, never on encoded strings. Only use hex and base64 for pretty-printing.
 */
func main() {
	data := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	str, err := hex.DecodeString(data)
	if err != nil {
		fmt.Printf("Cannot decode hex")
		os.Exit(1)
	}

	enc := base64.StdEncoding.EncodeToString(str)
	fmt.Println(enc)
}
