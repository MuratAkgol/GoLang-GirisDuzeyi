package defer_statement

import "fmt"

func A() {
	fmt.Println("a fonksiyonu çalıştı.")
}
func C() {
	fmt.Println("C fonksiyonu çalıştı.")
}
func D() {
	fmt.Println("d fonksiyonu çalıştı.")
}

func B() {
	defer A()
	defer C()
	defer D()
	fmt.Println("b fonksiyonu çalıştı")

}
