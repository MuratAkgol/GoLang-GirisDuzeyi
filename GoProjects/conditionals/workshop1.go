package conditionals

import "fmt"

func Demo1() {
	//üç adet int değişkeni tanımla.
	//Ekrana en büyük olanı yazdır.

	var sayi1, sayi2, sayi3 int = 10, 5, 30

	var enBuyuk int = sayi1

	if sayi2 > enBuyuk {
		enBuyuk = sayi2
	}
	if sayi3 > enBuyuk {
		enBuyuk = sayi3
	}

	fmt.Printf("En büyük sayı %v", enBuyuk)

}
