package cryptolib

import "encoding/hex"

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
