package main

import (
	"fmt"
	"slices"
)

func main() {
	result := revert("HEY")

	fmt.Printf("REVERT: %v\n", result)
}

func revert(s string) string {
	runes := []rune(s)

	slices.Reverse(runes)

	return string(runes)

}
