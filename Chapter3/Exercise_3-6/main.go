package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var alphabetChar = make([]rune, 0)
var plainText string = "I love working with Go!"
var cipherText = make([]rune, len(plainText))

// generateArray generates a slice of characters from
// 'A' to 'Z' and then shuffles them in order to randomize them.
func generateArray(b []rune) []rune {
	for i := 65; i >= 'A' && i <= 'Z'; i++ {
		alphabetChar = append(alphabetChar, rune(i))
	}
	rand.Shuffle(len(alphabetChar), func(i, j int) {
		alphabetChar[i], alphabetChar[j] = alphabetChar[j], alphabetChar[i]
	})
	return alphabetChar
}

// rot13 is a simple letter substitution cipher that replaces a letter
// with the 13th letter after it in the alphabet. It is used to
// encode/decode values passed into it.
func rot13(b rune) rune {
	var a, z rune
	if 'A' <= b && b <= 'Z' {
		a, z = 'A', 'Z'
	} else {
		return b
	}
	return (b-a+13)%(z-a+1) + a
}

// plainTextReader accepts a plaintext string, converts the string to
// uppercase letters and uses the rot13 function to return a string
// of encoded/decoded text.
func plainTextReader(s string) string {
	upperPlainText := strings.ToUpper(s)
	for i, v := range upperPlainText {
		cipherText[i] = rot13(v)
	}
	str := string(cipherText)
	return str
}

func main() {
	// First print the plainText string followed by the encoded cipher.
	fmt.Println(plainText)
	generateArray(alphabetChar)
	encodedString := plainTextReader(plainText)
	fmt.Println(encodedString)

	// Print the decoded string from the encoded cipher.
	decodedString := plainTextReader(string(cipherText))
	fmt.Println(decodedString)
}
