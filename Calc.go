// Calc
package main

import (
	"fmt"
)

func main() {
	fmt.Print("Простенький калькулятор.\n 1. Вводите первое число\n 2. Оператор\n 3. Второе число\n")
	fmt.Print("Допускаются следующие операторы:  +,  -,  *,  /  \n")
	fmt.Print(CalcFunc())
}
