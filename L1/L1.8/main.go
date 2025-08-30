package main

import "fmt"

func main() {
	revertBitAtIndex(int64(5), 0)
}

func revertBitAtIndex(word int64, index int) int64 {
	reversedWord := word ^ 1<<index

	fmt.Printf("After B: %0b\n", word)
	fmt.Printf("Before B: %0b\n", reversedWord)

	fmt.Printf("After D: %d\n", word)
	fmt.Printf("Before D: %d\n", reversedWord)

	return reversedWord
}
