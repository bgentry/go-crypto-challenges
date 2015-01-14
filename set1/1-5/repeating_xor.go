package repeating

import "github.com/bgentry/go-crypto-challenges/cryptolib"

func RepeatingXor(plaintext, shortkey []byte) []byte {
	key := cryptolib.GenerateRepeatingKey(shortkey, len(plaintext))
	return cryptolib.Xor(plaintext, key)
}
