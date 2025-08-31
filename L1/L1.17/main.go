package main

import "fmt"

func main() {
	result := binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8)

	fmt.Printf("RESULT: %v\n", result)
}

func binarySearch(slice []int, target int) int {
	left := 0
	right := len(slice) - 1
	mid := int((left + right) / 2)
	foundIndex := -1

	for left <= right {
		if target == slice[mid] {
			foundIndex = mid
			break
		} else if target > slice[mid] {
			left = mid + 1
			mid = int((left + right) / 2)

		} else {
			right = mid - 1
			mid = int((left + right) / 2)
		}

	}
	return foundIndex
}
