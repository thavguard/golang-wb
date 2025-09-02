package main

import (
	"fmt"

	currenttime "L2.8/current-time"
)

func main() {
	current := currenttime.GetCurrentTime()

	fmt.Printf("current: %v\n", current)
}
