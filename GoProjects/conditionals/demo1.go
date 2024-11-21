package conditionals

import "fmt"

func Conditonals() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen > hesap {
		fmt.Println("Bakiye yetersiz.")
	}
	if cekilmekIstenen < hesap {
		hesap = hesap - cekilmekIstenen
		fmt.Println("İşlem başarılı. Yeni bakiye : " + fmt.Sprintf("%v", hesap))
	}
}
