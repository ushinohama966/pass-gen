package internal

import "fmt"

func Gen(l uint16, excludeSpecialChar bool) {
	pass := generatePass(l, excludeSpecialChar)

	fmt.Println(pass) //nolint
}
