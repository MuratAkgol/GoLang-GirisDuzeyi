package variables

import "fmt"

func Demo1() {
	fmt.Println("Hello world!")

	var myFirstString string = "Selam!"
	fmt.Println(myFirstString)

	var myFirstInteger int = 23
	fmt.Println(myFirstInteger)

	diffrentSyntax := 25
	fmt.Println("Eğer aşağıdaki gibi yaparsan 2. parametrenin veri tipini yazdırabilirsin.")
	fmt.Printf("veri tipi : %T\n", diffrentSyntax)

	var metin1 string = "Murat"
	var metin2 string = "Akgöl"

	var durum bool = metin1 == metin2
	var durum2 bool = metin1 != metin2

	fmt.Println(durum, durum2)

}
