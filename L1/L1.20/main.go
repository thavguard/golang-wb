package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	result := reverse("snow dog sun")

	fmt.Printf("REVERSE: %v\n", result)
}

func reverse(s string) string {
	// Строки в go иммутабельны поэтому без слайсов сделать не получится.
	// Были мысли сделать через [strings.Builder], но он под капотом тоже использует слайс)
	// Еще можно просто путем конкатенации, но это очень бэк практик так как на каждую операцию создается новая строка что не есть хорошо
	slice := strings.Fields(s)

	slices.Reverse(slice)

	return strings.Join(slice, " ")
}
