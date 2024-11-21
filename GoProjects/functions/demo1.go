package functions

import "fmt"

func Sum(number1 int, number2 int) int {
	sum := number1 + number2
	return sum
}

func Minus(number1 int, number2 int) {
	fmt.Println("Sonuç:", number1-number2)
}
