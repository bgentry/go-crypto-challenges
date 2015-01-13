package cryptolib

import (
	"encoding/hex"
	"testing"
)

func TestXor(t *testing.T) {
	str1 := "1c0111001f010100061a024b53535009181c"
	str2 := "686974207468652062756c6c277320657965"
	buf1, err := hex.DecodeString(str1)
	if err != nil {
		t.Fatal(err)
	}
	buf2, err := hex.DecodeString(str2)
	if err != nil {
		t.Fatal(err)
	}

	res := Xor(buf1, buf2)
	got := hex.EncodeToString(res)
	want := "746865206b696420646f6e277420706c6179"
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
