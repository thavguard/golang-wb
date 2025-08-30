package main

import (
	"fmt"
)

func main() {
	var ch chan int

	AssertType(ch)
}

func AssertType(v any) {
	// fmt.Printf("TYPE IS %T\n", v) // все тоже самое но в 1 строку

	switch v.(type) {
	case int:
		fmt.Printf("TYPE IS int\n")

	case string:

		fmt.Printf("TYPE IS string\n")

	case bool:
		fmt.Printf("TYPE IS bool\n")

	case chan any, chan int, chan string, chan bool:
		fmt.Printf("TYPE IS chan\n")

	}
}
