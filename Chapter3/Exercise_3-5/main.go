package main

import (
	"fmt"
	"strings"
)

var alphabetChar = []rune{
	'A',
	'B',
	'C',
	'D',
	'E',
	'F',
	'G',
	'H',
	'I',
	'J',
	'K',
	'L',
	'M',
	'N',
	'O',
	'P',
	'Q',
	'R',
	'S',
	'T',
	'U',
	'V',
	'W',
	'X',
	'Y',
	'Z',
}

var plainText string = "I love working with Go!"
var cipherText = make([]rune, len(plainText))

func encoderDecoder(s string) string {
	upperPlainText := strings.ToUpper(s)
	for i, v := range upperPlainText {
		if v >= 'A' && v <= 'Z' {
			cipherText[i] = alphabetChar['Z'-v]
		} else {
			cipherText[i] = v
		}
	}
	str := string(cipherText)
	return str
}

func main() {
	// First print the plainText string followed by the encoded cipher.
	fmt.Println(plainText)
	encodedString := encoderDecoder(plainText)
	fmt.Println(encodedString)

	// Print the decoded string from the encoded cipher.
	decodedString := encoderDecoder(string(cipherText))
	fmt.Println(decodedString)
}
