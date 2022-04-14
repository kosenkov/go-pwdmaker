package pwdgenerator

import (
	"math/rand"
	"strings"
)

const (
	// list of lowercase letters
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"

	// list of uppercase letters
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// permitted digits
	Digits = "0123456789"

	// list of special symbols
	Symbols = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"

	AllCharSet = LowerLetters + UpperLetters + Digits + Symbols
)

const (
	MinLowerLetters = 2
	MinUpperLetters = 2
	MinDigits       = 2
	MinSymbols      = 1
)

// Temp function implamentation from https://golangbyexample.com/generate-random-password-golang
//
func Generate(salt string, lenght int, useSpecSymbols bool) string {
	var password strings.Builder

	// Set special character
	if useSpecSymbols {
		for i := 0; i < MinSymbols; i++ {
			random := rand.Intn(len(Symbols))
			password.WriteString(string(Symbols[random]))
		}
	}

	// Set numeric
	for i := 0; i < MinDigits; i++ {
		random := rand.Intn(len(Digits))
		password.WriteString(string(Digits[random]))
	}

	// Set uppercase
	for i := 0; i < MinUpperLetters; i++ {
		random := rand.Intn(len(UpperLetters))
		password.WriteString(string(UpperLetters[random]))
	}

	charSet := salt + AllCharSet
	remainingLength := lenght - MinSymbols - MinDigits - MinUpperLetters
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(charSet))
		password.WriteString(string(charSet[random]))
	}

	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})

	return string(inRune)
}
