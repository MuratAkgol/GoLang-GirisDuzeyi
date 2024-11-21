package main

import (
	"golesson/error_handling"
)

func main() {
	//variables.Demo1()
	//conditionals.Conditonals()
	//conditonals.Demo2()
	//conditonals.Demo1()
	//loops.Demo1()
	//loops.Demo2()
	//loops.Demo3()
	//loops.Demo4()
	//arrays.Demo1()
	//arrays.Demo2()
	//arrays.Demo3()
	//arrays.Demo4()
	//slices.Demo1()
	//slices.Demo2()
	// var answer = functions.Sum(2, 6)
	// fmt.Println(answer)

	// bu şekilde parametreye _ koyarsan o parametre return olmaz.
	//answer1, answer2, answer3, _ := functions.MultipleReturnFunction(10, 6)

	// answer1, answer2, answer3, answer4 := functions.MultipleReturnFunction(10, 6)

	// fmt.Println("Toplam:", answer1)
	// fmt.Println("Çıkarma:", answer2)
	// fmt.Println("Çarpma:", answer3)
	// fmt.Println("Bölme:", answer4)

	//fmt.Println(functions.SumVariadic(1, 2, 3, 4))

	// numbers := []int{1, 2, 3, 4}
	// fmt.Println(functions.SumVariadic(numbers...))
	//maps.Demo1()
	//ranges.Demo1()
	//ranges.Demo2()
	//ranges.Demo3()

	// numb := 20
	// pointers.Demo1(&numb)
	// fmt.Println(numb)

	// numbers := []int{1, 2, 3}
	// pointers.Demo2(numbers)
	// fmt.Println("Maindeki sayı:", numbers[0])
	//structs.Demo1()
	//structs.Demo2()

	// TotalEvenNumbersCn := make(chan int)
	// TotalOddNumbersCn := make(chan int)
	// go channels.EvenNumbers(TotalEvenNumbersCn)
	// go channels.OddNumbers(TotalOddNumbersCn)
	// var carpim int = <-TotalEvenNumbersCn * <-TotalOddNumbersCn
	// fmt.Println("Çarpım", carpim)

	//interfaces.Demo1()
	// interfaces.Demo2()
	//interfaces.Demo3()
	// defer_statement.B()
	//defer_statement.Demo2_test()
	//defer_statement.Demo3()
	//error_handling.Demo1()
	error_handling.Demo2()
}
