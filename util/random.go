package util

import "math/rand"

func RandomString(n int) string {
	var key = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	res := ""
	for i := 0; i < n; i++ {
		res += string(key[rand.Intn(len(key))])
	}
	return res
}
