package decrypt_aes_ecb

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestDecryptAESECB(t *testing.T) {
	cipher, err := ioutil.ReadFile("7.txt")
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

	key := []byte("YELLOW SUBMARINE")
	plaintext, err := DecryptAESECB(cipherb, key)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("recovered plaintext:\n==========================")
	fmt.Printf("%s\n", plaintext)
}
