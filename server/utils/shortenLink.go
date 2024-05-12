package utils

import "math/rand"

func GenereateShortKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6

	shortkey := make([]byte, keyLength)

	for i := range shortkey{
		shortkey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortkey)
}