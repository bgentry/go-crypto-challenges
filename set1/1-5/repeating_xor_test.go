package repeating

import (
	"encoding/hex"
	"testing"
)

var plaintext = `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`

func TestRepeatingXor(t *testing.T) {
	wanthex := `0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f`
	got := RepeatingXor([]byte(plaintext), []byte("ICE"))
	gothex := hex.EncodeToString(got)

	if wanthex != gothex {
		t.Errorf("want %s, got %s", wanthex, gothex)
	}
}
