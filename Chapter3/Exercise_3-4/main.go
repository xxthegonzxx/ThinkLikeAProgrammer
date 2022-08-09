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

func main() {
	plainText := "I love working with Go!"
	upperPlainText := strings.ToUpper(plainText)
	cipherText := make([]rune, len(plainText))
	for i, v := range upperPlainText {
		if v >= 'A' && v <= 'Z' {
			cipherText[i] = alphabetChar['Z'-v]
		} else {
			cipherText[i] = v
		}
	}
	fmt.Println(plainText)
	fmt.Println(string(cipherText))
}
