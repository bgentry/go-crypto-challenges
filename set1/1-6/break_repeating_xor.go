package breakrepeating

import (
	"fmt"

	"github.com/bgentry/go-crypto-challenges/cryptolib"
)

func BreakRepeatingKeyXor(cipher []byte, minKeySize, maxKeySize int) (key []byte, plaintext string, err error) {
	likelyKeySize, err := findBestKeySize(cipher, minKeySize, maxKeySize)
	if err != nil {
		return nil, "", err
	}

	key = make([]byte, likelyKeySize)
	blocks := breakIntoBlocks(cipher, likelyKeySize)
	transposed := transposeBlocks(blocks)
	for i, t := range transposed {
		_, c, _ := cryptolib.FindBestPlaintext(t)
		key[i] = c
	}

	repeatedKey := cryptolib.GenerateRepeatingKey(key, len(cipher))
	plaintext = string(cryptolib.Xor(cipher, repeatedKey))

	return key, plaintext, nil
}

func breakIntoBlocks(data []byte, size int) [][]byte {
	blockCount := len(data) / size
	blocks := make([][]byte, blockCount)
	for i := 0; i < blockCount; i++ {
		blocks[i] = data[i*size : (i+1)*size]
	}
	return blocks
}

func transposeBlocks(blocks [][]byte) [][]byte {
	if len(blocks) == 0 {
		panic("WTF")
	}
	transposed := make([][]byte, len(blocks[0]))
	for i := range transposed {
		transposed[i] = make([]byte, len(blocks))
		for j := range blocks {
			transposed[i][j] = blocks[j][i]
		}
	}
	return transposed
}

func findBestKeySize(data []byte, minKeySize, maxKeySize int) (int, error) {
	if maxKeySize < minKeySize || minKeySize == 0 {
		panic("WTF")
	}

	bestKeySize := 0
	bestDist := float32(9999999) // higher than any actual distance

	for keySize := minKeySize; keySize <= maxKeySize; keySize++ {
		avgDist := float32(0)
		numIterations := 0

		for offset := 0; offset+2*keySize < len(data); offset += keySize {
			if len(data) < 2*keySize {
				return 0, fmt.Errorf("data wasn't long enough to compare two slices w/ key size %d", keySize)
			}
			firsti, secondi := data[offset:offset+keySize], data[offset+keySize:offset+2*keySize]
			dist, err := cryptolib.HammingDistance(firsti, secondi)
			if err != nil {
				return 0, err
			}
			avgDist += float32(dist) / float32(keySize)
			numIterations++
		}
		dist := avgDist / float32(numIterations)

		if dist < bestDist {
			bestDist = dist
			bestKeySize = keySize
		}
	}

	return bestKeySize, nil
}
