package main

import (
	"fmt"

	"github.com/bgentry/go-crypto-challenges/cryptolib"
)

func main() {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	inb := cryptolib.MustDecodeHex(input)
	fmt.Printf("input: %q\n", string(inb))

	var (
		highestScore       float32
		highestScoreChar   byte
		highestScoreString string
	)
	for i := 0; i < 255; i++ {
		keyByte := byte(i)
		key := cryptolib.GenerateUniformKey(keyByte, len(inb))
		got := string(cryptolib.Xor(inb, key))
		score := cryptolib.ScoreEnglish(got)
		if score > highestScore {
			highestScore = score
			highestScoreChar = keyByte
			highestScoreString = got
		}
	}
	fmt.Printf("highest score was %f for char %q. Plaintext was likely:\n%q\n",
		highestScore,
		highestScoreChar,
		highestScoreString,
	)
}
