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
