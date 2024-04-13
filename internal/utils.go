package internal

import (
	"crypto/rand"
	"math/big"
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

func generatePass(length uint16, excludeSpecialChar bool) string {
	letters := letterJoin(
		getUppercaseCharacters(),
		getLowercaseCharacters(),
		getNumberCharacters(),
	)
	if !excludeSpecialChar {
		letters = append(letters, getSpecialCharacters()...)
	}
	pass := make([]byte, 0, length)
	for range length {
		pass = append(pass, generateChar(letters))
	}

	return string(pass)
}
