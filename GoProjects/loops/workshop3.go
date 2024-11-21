package loops

import "fmt"

func Demo4() {
	sayi1 := 1184
	countForSayi1 := 0

	sayi2 := 1210
	countForSayi2 := 0

	for i := 1; i < sayi1; i++ {
		if sayi1%i == 0 {
			countForSayi1 += i
		}
	}
	for i := 1; i < sayi2; i++ {
		if sayi2%i == 0 {
			countForSayi2 += i
		}
	}

	if countForSayi1 == sayi2 && countForSayi2 == sayi1 {
		fmt.Printf("%v ve %v 'Arkadaş' sayıdır.", sayi1, sayi2)
	} else {
		fmt.Printf("%v ve %v 'Arkadaş' sayı değildir.", sayi1, sayi2)
	}
}
