package main

import (
	"fmt"
	"math/big"
)

func main() {
	bigA := new(big.Int)
	bigA.SetString("24000000000000000000000000000000", 10)

	bigB := new(big.Int)
	bigB.SetString("12000000000000000000000000000000", 10)

	// Деление
	divide := new(big.Int)
	divide.Div(bigA, bigB)
	fmt.Printf("divide: %v\n", divide)

	// Умножение
	mult := new(big.Int)
	mult.Mul(bigA, bigB)
	fmt.Printf("mult: %v\n", mult)

	// Сложение
	plus := new(big.Int)
	plus.Add(bigA, bigB)
	fmt.Printf("plus: %v\n", plus)

	// Вычитание
	minus := new(big.Int)
	minus.Sub(bigA, bigB)
	fmt.Printf("minus: %v\n", minus)

}
