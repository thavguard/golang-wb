package main

import (
	"fmt"
	"time"

	"L1.25/cat"
	mysleep "L1.25/mySleep"
)

func main() {
	fmt.Printf("Через 4 секунды появится котик\n")

	mysleep.MySleepFor(2 * time.Second)
	mysleep.MySleep(2 * time.Second)

	cat.Run()
}
