package breakrepeating

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestBreakRepeatingKeyXor(t *testing.T) {
	cipher, err := ioutil.ReadFile("6.txt")
	if err != nil {
		t.Fatal(err)
	}
	var cipherb []byte
	for _, line := range bytes.Split(cipher, []byte{'\n'}) {
		data, err := base64.StdEncoding.DecodeString(string(line))
		if err != nil {
			t.Fatal(err)
		}
		cipherb = append(cipherb, data...)
	}

	key, plaintext, err := BreakRepeatingKeyXor(cipherb, 2, 60)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("cracked key: %x\n", key)
	fmt.Println("recovered plaintext:\n==========================")
	fmt.Println(plaintext)
}
