package functions

func SumVariadic(numbers ...int) int {
	sum := 0

	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}
