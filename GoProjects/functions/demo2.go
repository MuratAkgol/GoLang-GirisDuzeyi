package functions

func MultipleReturnFunction(number1 int, number2 int) (int, int, int, float32) {
	sum := number1 + number2
	minus := number1 - number2
	multiply := number1 * number2
	divide := float32(number1 / number2)

	return sum, minus, multiply, divide
}
