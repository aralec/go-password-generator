package gopasswordgenerator

import (
	"math/rand"
	"strings"
)

// randomize returns a randomized string array from source
func randomize(source string, size uint) []string {
	passwords := make([]string, 0)
	for i := 0; i < int(size); i++ {
		passwords = append(passwords, string(source[rand.Intn(len(source))]))
	}
	return passwords
}

// mixCase randomizes the case of the inputed string array
func mixCase(s []string) {
	for i := range s {
		if rand.Intn(2) == 1 {
			s[i] = strings.ToUpper(s[i])
		} else {
			s[i] = strings.ToLower(s[i])
		}
	}
}
