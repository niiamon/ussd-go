package ussd

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
)

var alphaNumeric = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// StrLower returns lowercase s.
func StrLower(s string) string {
	return strings.ToLower(s)
}

// StrTrim returns s with excess whitespace removed.
func StrTrim(s string) string {
	return strings.TrimSpace(s)
}

// StrRandom returns random string of given length.
func StrRandom(length int) string {
	result := make([]rune, length)
	for i := range result {
		result[i] = alphaNumeric[rand.Intn(len(alphaNumeric))]
	}
	return string(result)
}

// panicln takes formatted string and panics with an error.
func panicln(format string, a ...interface{}) {
	log.Panicln(fmt.Errorf(format, a...))
}
