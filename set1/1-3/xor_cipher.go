package main

import (
	"fmt"

	"github.com/bgentry/go-crypto-challenges/cryptolib"
)

func main() {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	inb := cryptolib.MustDecodeHex(input)
	fmt.Printf("input: %q\n", string(inb))

	score, char, plaintext := cryptolib.FindBestPlaintext(inb)
	fmt.Printf("highest score was %f for char %q. Plaintext was likely:\n%q\n",
		score,
		char,
		plaintext,
	)
}
