package main

import (
	"fmt"

	"L2.9/unpackstr"
)

func main() {
	input := `s1`
	test := ""
	result, err := unpackstr.Unpack(input)

	if err != nil {
		fmt.Printf("ERR: %v\n", err)
		return
	}

	fmt.Printf("input  : %v\n", input)
	fmt.Printf("result : %v\n", result)
	fmt.Printf("test   : %v\n", test)

}
