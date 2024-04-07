package internal

import (
	"crypto/rand"
	"math/big"
)

const (
	defaultPassLength = 12
)

func generateChar(letters []byte) byte {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))

	return letters[n.Uint64()]
}

func letterJoin(letters ...[]byte) []byte {
	letter := make([]byte, 0)
	for _, l := range letters {
		letter = append(letter, l...)
	}

	return letter
}

func generatePass() string {
	letters := letterJoin(
		uppercaseCharacters,
		lowercaseCharacters,
		numberCharacters,
		specialCharacters,
	)
	pass := make([]byte, 0, defaultPassLength)
	for range defaultPassLength {
		pass = append(pass, generateChar(letters))
	}

	return string(pass)
}
