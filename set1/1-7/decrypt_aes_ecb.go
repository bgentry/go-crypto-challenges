package decrypt_aes_ecb

import "crypto/aes"

func DecryptAESECB(ciphertext, key []byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	bs := block.BlockSize()
	if len(ciphertext)%bs != 0 {
		panic("Need a multiple of the blocksize")
	}

	plaintext = make([]byte, len(ciphertext))
	ptblock := plaintext
	for len(ciphertext) > 0 {
		block.Decrypt(ptblock, ciphertext[:bs])
		ptblock = ptblock[bs:]
		ciphertext = ciphertext[bs:]
	}
	return plaintext, nil
}
