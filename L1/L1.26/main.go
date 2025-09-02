package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Unique"

	isUnique := checkSring(s)

	fmt.Printf("isUnique: %v\n", isUnique)
}

func checkSring(s string) bool {
	myset := make(map[rune]struct{})

	for _, c := range strings.ToLower(s) {
		myset[c] = struct{}{}
	}

	return len([]rune(s)) == len(myset)
}
