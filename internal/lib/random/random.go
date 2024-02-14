package random

import (
	"math/rand/v2"
)

func NewRandomString(l int) string {
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789")

	s := make([]rune, l)

	for i := range s {
		s[i] = letters[rand.IntN(len(letters))]
	}

	return string(s)
}
