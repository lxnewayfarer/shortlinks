package lib

import (
	"math/rand"
	"os"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func ShortenLink(link string) (string, error) {
	appUrl := os.Getenv("APP_URL")

	// check cache
	// return if cache
	// randSeq while exists randSeq in cache
	// save link
	// save cache

	return appUrl + "/" + randSeq(8), nil
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
