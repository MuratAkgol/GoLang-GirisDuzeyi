package loops

import (
	"fmt"
	"math/rand"
)

func Demo2() {

	aklimdakiSayi := rand.Intn(20)
	//aklimdakiSayi := 80
	tahminEdilenSayi := 0

	fmt.Println("Bir sayı tahmin ediniz.")
	fmt.Scanln(&tahminEdilenSayi)

	//istersek şöyle bir for yazılabilir.
	// for aklimdakiSayi !=tahminEdilenSayi{
	// 	//kodlar burada olacak.
	// }

	for i := 0; i < 5; i++ {
		fmt.Println("Kalan hak:" + fmt.Sprintf("%v", 5-i))

		if tahminEdilenSayi < aklimdakiSayi {
			fmt.Println("Daha yüksek bir sayı")
			fmt.Scanln(&tahminEdilenSayi)
		} else if aklimdakiSayi < tahminEdilenSayi {
			fmt.Println("Daha düşük bir sayı")
			fmt.Scanln(&tahminEdilenSayi)
		} else {
			fmt.Println("Sayıyı bildiniz.")
			break
		}

		if i == 5 {
			fmt.Println("Hakkınız bitti.")
			break
		}
	}

}
