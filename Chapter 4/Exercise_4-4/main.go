package main

import "fmt"

type arrayString []byte

func (a *arrayString) append(r byte) {
	*a = append((*a), r)
	(*a)[0]++
}

func (a arrayString) charAt(index int) byte {
	return a[index+1]
}

func (a *arrayString) concatenate(b arrayString) {
	*a = append(*a, b[1:]...)
	(*a)[0] = byte(int((*a)[0]) + len(b) - 1)
}

func (a arrayString) output() string {
	return string(a[1:])
}

func printOutput(a arrayString) {
	fmt.Println("size of array:", a[0])
	fmt.Println((&a).output())
	fmt.Println(string((&a).charAt(0)))
}

func main() {
	a := arrayString{byte(5), 'g', 'o', 'l', 'a', 'n', 'g'}
	(&a).append('!')
	printOutput(a)

	(&a).concatenate(a)
	printOutput(a)
}
