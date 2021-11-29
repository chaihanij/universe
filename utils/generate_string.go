package utils

import "math/rand"

func GenerateRandomString(strLength int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	strGenerate := make([]byte, strLength)
	for i := range strGenerate {
		strGenerate[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(strGenerate)
}
