package conditionals

import "fmt"

func Demo2() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen > hesap {
		fmt.Println("Bakiye yetersiz.")
	} else if cekilmekIstenen == hesap {
		fmt.Println("Bakiyeniz sıfırlandı.")
	} else {
		hesap = hesap - cekilmekIstenen
		fmt.Println("İşlem başarılı. Yeni bakiye : " + fmt.Sprintf("%v", hesap))
	}
}
