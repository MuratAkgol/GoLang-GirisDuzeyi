package pointers

import "fmt"

func Demo2(numbers []int) {
	numbers[0] = 100
	fmt.Println("Demodaki sayı:", numbers[0])
}
