package main

import (
	"fmt"
)

type arrayString []string

func (a *arrayString) replaceString(source, target, replaceText arrayString) {
	sourceLength := len(source)
	replaceIndexes := make([]int, 0)
	targetLength := len(target)

	// special case: empty arrayString as both source and target
	if sourceLength == 0 && targetLength == 0 {
		*a = replaceText
		return
	}

Outer:
	for i := 0; i < sourceLength; i++ {
		for j := 0; j < targetLength; j++ {
			outOfBounds := i+j >= sourceLength
			mismatch := source[i+j] != target[j]

			if outOfBounds || mismatch {
				continue Outer
			}
		}
		replaceIndexes = append([]int{i}, replaceIndexes...)
	}

	for _, i := range replaceIndexes {
		end := append(arrayString{}, source[i+targetLength:]...)

		source = append(source[:i], replaceText...)
		source = append(source, end...)
	}

	*a = source
}

func main() {
	a := arrayString{"a", "b", "c", "d", "a", "b", "e", "e"}
	b := arrayString{"a", "b"}
	replaceStr := arrayString{"x", "y", "z"}
	a.replaceString(a, b, replaceStr)
	fmt.Println(a)
}
