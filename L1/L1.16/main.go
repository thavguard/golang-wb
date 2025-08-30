package main

import "fmt"

func main() {
	integers := []int{5, 7, 1, 9, 10, 43, 65, 3, 34, 24, 32, 4, 234, 23, 42, 22, 34, 645, 7, 6, 7, 56, 7}
	result := quickSort(integers)

	fmt.Printf("RESULT: %v\n", result)
}

func quickSort(slice []int) []int {

	if len(slice) < 1 {
		return slice
	}

	pivot := slice[int(len(slice)/2)]

	left := []int{}
	right := []int{}
	equal := []int{}

	for _, item := range slice {
		if item > pivot {
			right = append(right, item)
		} else if item < pivot {
			left = append(left, item)
		} else {
			equal = append(equal, item)
		}
	}

	result := []int{}

	result = append(result, quickSort(left)...)
	result = append(result, equal...)
	result = append(result, quickSort(right)...)

	return result

}
