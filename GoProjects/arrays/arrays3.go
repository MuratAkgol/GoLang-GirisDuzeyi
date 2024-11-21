package arrays

import "fmt"

func Demo3() {
	numbers := [5]int{20, 30, 40, 10, 2}
	biggestNumber := numbers[0]
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > biggestNumber {
			biggestNumber = numbers[i]
		}
	}
	fmt.Println("En büyük sayı:", biggestNumber)

}
