package defer_statement

import "fmt"

func Demo2(numbers int) string {
	defer DeferFunc()
	if numbers%2 == 0 {
		fmt.Println("Çift sayı çalıştı")
		return "Çift Sayıdır."
	} else {
		fmt.Println("Tek sayı çalıştı")
		return "Tek sayıdır."
	}

}

func Demo2_test() {
	fmt.Println(Demo2(9))
}

func DeferFunc() {
	fmt.Println("Bitti.")
}
