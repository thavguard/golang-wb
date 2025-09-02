package main

import (
	"fmt"
	"reflect"
)

func main() {
	slice := []string{"Удалить", "i-ый", "элемент", "из", "слайса"}

	removeFromSlice(&slice, 1)

	fmt.Printf("result: %v\n", slice)
	fmt.Printf("len(slice): %v\n", len(slice))
	fmt.Printf("cap(slice): %v\n", cap(slice))

}

func removeFromSlice(slice *[]string, index int) {

	copy((*slice)[index:], (*slice)[index+1:])

	val := reflect.Indirect(reflect.ValueOf(slice))

	(*slice)[len(*slice)-1] = ""

	val.SetLen(len(*slice) - 1)
	val.SetCap(len(*slice))

}
