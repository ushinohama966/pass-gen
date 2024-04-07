package internal

import "fmt"

func Gen(l uint16) {
	pass := generatePass(l)

	fmt.Println(pass) //nolint
}
