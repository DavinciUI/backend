package code

import (
	"math/rand"
	"strings"
)

// GenerateNewCode create a random new code of 10 digits
func GenerateNewCode() *Code {
	return &Code{randomCode(10)}
}

type Code struct {
	Code string
}

func (code Code) String() string {
	return code.Code
}

var combinations = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randomCode(size int) string {
	codeBuilder := strings.Builder{}

	for i := 0; i < size; i++ {
		codeBuilder.WriteString(string(combinations[rand.Intn(len(combinations))]))
	}

	return codeBuilder.String()
}