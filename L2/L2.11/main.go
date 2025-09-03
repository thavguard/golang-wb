package main

import (
	"fmt"
	"slices"
)

func main() {
	input := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}

	m := make(map[string][]string)
	annMap := make(map[string]string) // мапа чтобы хранить анограмму для каждого значения чтобы не сортировать ее постоянно

	for _, i := range input {
		currentAnogram := toSortedBytesSlice(i)

		if hasAnogram, ok := annMap[currentAnogram]; !ok {
			annMap[currentAnogram] = i
		} else {
			m[hasAnogram] = append(m[hasAnogram], i)
		}

	}

	for k, v := range m {
		m[k] = append(v, k)
	}

	fmt.Printf("annMap: %v\n", len(annMap))
	fmt.Printf("m: %v\n", m)

}

func toSortedBytesSlice(s string) string {
	slice := []byte(s)

	slices.Sort(slice)

	return string(slice)

}
