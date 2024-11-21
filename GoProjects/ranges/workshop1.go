package ranges

import "fmt"

//1-10 arasındaki sayılardan tek olanları toplayan program
func Demo2() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	total := 0
	for _, v := range numbers {
		if v%2 != 0 {
			total += v
		}
	}
	fmt.Println("Tek sayıların toplamı:", total)
}
