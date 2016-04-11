package main

import (
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	var str1, str2 string
	fmt.Println("Enter first hex string")
	_, err := fmt.Scanf("%s", &str1)
	if err != nil {
		fmt.Println("Cannot read string")
		os.Exit(1)
	}

	fmt.Println("Enter second hex string")
	_, err = fmt.Scanf("%s", &str2)
	if err != nil {
		fmt.Println("Cannot read string")
		os.Exit(1)
	}

	// Check if both the strings are of equal length
	if len(str1) != len(str2) {
		fmt.Println("Strings must be of equal lengths")
		os.Exit(1)
	}

	h1, er := hex.DecodeString(str1)
	if er != nil {
		fmt.Println("Not a valid hex string")
		os.Exit(1)
	}

	h2, er := hex.DecodeString(str2)
	if er != nil {
		fmt.Println("Not a valid hex string")
		os.Exit(1)
	}

	fmt.Printf("%x\n", getXOR(h1, h2))
}

// getgetXOR returns the XOR of two strings
func getXOR(s1, s2 []byte) string {
	b := make([]byte, len(s1))
	for i := 0; i < len(s1); i++ {
		b[i] = s1[i] ^ s2[i]
	}

	return string(b)
}
