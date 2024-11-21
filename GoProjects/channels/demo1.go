package channels

func EvenNumbers(evenNumberCn chan int) {
	toplam := 0
	for i := 0; i <= 10; i += 2 {
		toplam += i
	}
	// Channel'a bu şekilde data atılır.
	evenNumberCn <- toplam
}

func OddNumbers(oddNumberCn chan int) {
	toplam := 0
	for i := 1; i <= 10; i += 2 {
		toplam += i
	}

	oddNumberCn <- toplam
}
