package rand

import (
	"math/rand"
	"time"
)

var (
	chars = "1234567890QWERTYUIOPMLKJHGFDSAZXCVBNqwertyuiopmlkjhgfdsazxcvbn"
)

func RandBytes(length int) []byte {
	rand.Seed(time.Now().UTC().UnixNano())
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return result
}

func RandString(length int) string {
	return string(RandBytes(length))
}
