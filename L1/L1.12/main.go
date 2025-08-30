package main

import (
	"fmt"

	myset "L1.12/my-set"
)

func main() {
	set := myset.NewMySet()

	set.Append("cat", "cat", "dog", "cat", "tree")

	fmt.Printf("%v\n", set)

}
