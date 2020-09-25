package utils

import (
	"math/rand"
	"time"
)

func RandCode(randString string,length int) string{
	var letters = []rune(randString)
	rand.Seed(time.Now().Unix())
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

