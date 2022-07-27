package main

import (
	"fmt"
)

type modeType int

const (
	Lowercase modeType = iota
	Uppercase
	Punctuation
)

func main() {
	var digitChar byte
	var outputCharacter byte
	var mode modeType = Uppercase
	for {
		_, err := fmt.Scanf("%c", &digitChar)
		if err != nil {
			fmt.Println("Error:", err)
		}
		number := int(digitChar - '0')
		fmt.Scanf("%c", &digitChar)
		for (digitChar != 10) && (digitChar != ',') {
			number = (number*10 + (int(digitChar) - '0'))
			fmt.Scanf("%c", &digitChar)
		}
		switch mode {
		case Uppercase:
			number = number % 27
			outputCharacter = byte(number) + 'A' - 1
			if number == 0 {
				mode = Lowercase
				continue
			}
			break
		case Lowercase:
			number = number % 27
			outputCharacter = byte(number) + 'a' - 1
			if number == 0 {
				mode = Punctuation
				continue
			}
			break
		case Punctuation:
			number = number % 9
			switch number {
			case 1:
				outputCharacter = '!'
			case 2:
				outputCharacter = '?'
			case 3:
				outputCharacter = ','
			case 4:
				outputCharacter = '.'
			case 5:
				outputCharacter = ' '
			case 6:
				outputCharacter = ';'
			case 7:
				outputCharacter = '"'
			case 8:
				outputCharacter = '\''
			}
			if number == 0 {
				mode = Uppercase
				continue
			}
			break
		}
		data := []byte{outputCharacter}
		fmt.Println(string(data))
		if !(digitChar != 10) {
			break
		}
	}
}
