package main

import "testing"

func TestHexToBase64(t *testing.T) {
	b64, err := hexToBase64(hexString)
	if err != nil {
		t.Fatal(err)
	}
	want := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	if b64 != want {
		t.Errorf("want %q, got %q", want, b64)
	}
}
