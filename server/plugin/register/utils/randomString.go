package utils

import "math/rand"

func RandomString(length int) (reStr string) {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i:=0; i < length; i++ {
		reStr += string(str[rand.Intn(len(str))])
	}
	return reStr
}