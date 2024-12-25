package gopasswordgenerator

import (
	"math/rand"
	"strings"
)

var (
	lowerCaseAsciiCharactersString string = "abcdefghijklmnopqrstuvwxyz"
	digitsString                          = "0123456789"
	specialCharactersString               = "!@#$%&*_?-"
)

// Genenerate generates a random password of the given length.
func Generate(desiredLength uint) (string, error) {
	if desiredLength < 8 {
		return "", &PasswordSizeError{}
	}

	// Getting ratios of letters, numbers and special characters :
	lr := desiredLength / 2
	nr := desiredLength - lr
	sr := nr

	// Getting random sets:
	letters := randomize(lowerCaseAsciiCharactersString, lr)
	numbers := randomize(digitsString, nr)
	specialCharacters := randomize(specialCharactersString, sr)

	// From all-lowercase to randomly mixed case array of letters :
	mixCase(letters)

	password := append(letters, numbers...)
	password = append(password, specialCharacters...)

	// After appending numbers and special characters to letters, shuffle the array :
	rand.Shuffle(len(password), func(i, j int) {
		if i != j {
			password[i], password[j] = password[j], password[i]
		}
	})

	return strings.Join(password, ""), nil
}
