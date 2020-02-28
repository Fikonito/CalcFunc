// CalcFunc
package main

import (
	"fmt"
)

func CalcFunc() int {
	var a, b, sum int
	op := ""

	fmt.Print("Введите первое число: \n")
	fmt.Scan(&a)

	fmt.Print("Введите оператор: \n")
	fmt.Scan(&op)

	fmt.Print("Введите второе число: \n")
	fmt.Scan(&b)

	switch op {
	case "+":
		sum = a + b
	case "-":
		sum = a - b
	case "*":
		sum = a * b
	case "/":
		sum = a / b
	default:
		fmt.Printf("Вы ввели %s. Такого оператора не существует", op)
	}
	return sum
}
