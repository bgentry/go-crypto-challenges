package cryptolib

import (
	"bytes"
	"encoding/hex"
	"errors"
)

func HammingDistance(a, b []byte) (diff int, err error) {
	if len(a) != len(b) {
		return 0, errors.New("length mismatch")
	}
	xor := Xor(a, b)
	for i := range xor {
		diff += int(popcnt[xor[i]])
	}
	return
}

// Set bits count in a byte
var popcnt = [256]byte{
	0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, // 0
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, // 1
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, // 2
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, // 3
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, // 4
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, // 5
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, // 6
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, // 7
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, // 8
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, // 9
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, // 10
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, // 11
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, // 12
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, // 13
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, // 14
	4, 5, 5, 6, 5, 6, 6, 7, 5, 6, 6, 7, 6, 7, 7, 8, // 15
}

func MustDecodeHex(hexstr string) []byte {
	inb, err := hex.DecodeString(hexstr)
	if err != nil {
		panic(err)
	}
	return inb
}

func Xor(buf1, buf2 []byte) []byte {
	res := make([]byte, len(buf1))
	for i := range buf1 {
		res[i] = buf1[i] ^ buf2[i]
	}
	return res
}

func GenerateRepeatingKey(k []byte, length int) []byte {
	n := (length + len(k) - 1) / len(k)
	key := bytes.Repeat(k, n)
	return key[:length]
}

func GenerateUniformKey(c byte, length int) []byte {
	key := make([]byte, length)
	for j := range key {
		key[j] = byte(c)
	}
	return key
}

// EnglishLetterFrequencies contains the relative frequencies (percentages) of
// letters in the English lanuage.
//
// Source:
// https://en.wikipedia.org/wiki/Letter_frequency#Relative_frequencies_of_letters_in_the_English_language
var EnglishLetterFrequencies = map[string]float32{
	"a": 8.167, "b": 1.492, "c": 2.782, "d": 4.253, "e": 12.702,
	"f": 2.228, "g": 2.015, "h": 6.094, "i": 6.966, "j": 0.153,
	"k": 0.772, "l": 4.025, "m": 2.406, "n": 6.749, "o": 7.507,
	"p": 1.929, "q": 0.095, "r": 5.987, "s": 6.327, "t": 9.056,
	"u": 2.758, "v": 2.360, "x": 0.150, "y": 1.974, "z": 0.074,
	" ": 13.0, // space is slightly more frequent than (e)
}

func ScoreEnglish(plaintext string) (score float32) {
	for _, char := range plaintext {
		score += EnglishLetterFrequencies[string(char)]
	}
	return
}

func FindBestPlaintext(cipher []byte) (highestScore float32, char byte, plaintext string) {
	keySize := len(cipher)
	for i := 0; i < 255; i++ {
		thisChar := byte(i)
		thisKey := GenerateUniformKey(thisChar, keySize)
		thisPlain := string(Xor(cipher, thisKey))
		score := ScoreEnglish(thisPlain)
		if score > highestScore {
			highestScore = score
			char = thisChar
			plaintext = thisPlain
		}
	}
	return
}
