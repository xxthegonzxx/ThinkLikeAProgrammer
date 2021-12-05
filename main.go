package main

import (
	"fmt"
	"log"
)

func main() {

	fmt.Println("Enter a six digit number, 0-9: ")
	fmt.Println("---------------------")
	var digit int
	fmt.Print("-> ")
	_, err := fmt.Scanf("%d", &digit)
	if err != nil {
		log.Fatal(err)
	}
	// doubledDigitValue(&digit)
	var bt byte
	bt = byte(digit)
	readIndividualNumber(bt)
}

func doubledDigitValue(n *int) int {
	var temp int
	temp = *n
	var doubledDigit = temp * 2
	var sum int
	if doubledDigit >= 10 {
		sum = 1 + doubledDigit%10
	} else {
		sum = doubledDigit
	}
	fmt.Println(sum)
	return sum
}

func readIndividualNumber(m byte) {
	checksum := byte(0)
	position := 1
	for position != 10 {
		if position%2 == 0 {
			checksum += m - byte(0)
		} else {
			asInt := int(m - byte(0))
			checksum += byte(doubledDigitValue(&asInt))
		}
		position++
	}
	fmt.Println("Checksum is:", checksum)
	if checksum%10 == 0 {
		fmt.Println("Checksum is divisble by 10. Valid.")
	} else {
		fmt.Println("Checksum is not divisible by 10. Invalid.")
	}
}
