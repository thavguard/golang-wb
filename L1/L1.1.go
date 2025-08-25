package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (human *Human) Meet() {
	fmt.Printf("Hello World! I`m %v\n", human.Name)
}

type Action struct {
	speed int
	Human
}

func (act *Action) Walk() {
	act.speed = act.speed + 1

	act.Meet() // method of Human struct
}

func main() {

}
