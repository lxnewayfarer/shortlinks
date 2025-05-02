package lib

import (
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type Random interface {
	RandSeq(n int) string
}

type RandomInstance struct {}
type MockRandomInstance struct {
	Value string
}

func (r RandomInstance) RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (r MockRandomInstance) RandSeq(n int) string {
	return r.Value
}
