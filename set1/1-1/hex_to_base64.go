package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
)

func hexToBase64(hexstr string) (string, error) {
	b, err := hex.DecodeString(hexstr)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

const hexString = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

func main() {
	b64, err := hexToBase64(hexString)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(b64)
}
