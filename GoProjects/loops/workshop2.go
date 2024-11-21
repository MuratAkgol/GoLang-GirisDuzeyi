package loops

import "fmt"

func Demo3() {
	//kullanıcıdan bir sayı girmesini iste.
	//asal olup olmadığını bulan bir örnek yaz
	sayi := 0
	count := 0
	fmt.Println("Bir sayı giriniz:")
	fmt.Scanln(&sayi)

	for i := 1; i < sayi; i++ {
		if sayi%i == 0 {
			count++
		}
	}
	if count == 1 {
		fmt.Println("Girilen sayı asaldır.")
	} else {
		fmt.Println("Girilen sayı asal değildir.")
	}
}
