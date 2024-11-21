package structs

import "fmt"

func Demo1() {
	fmt.Println(product{name: "Computer", price: 40000, brand: "XYZ", discountRate: 20})
}

type product struct {
	name         string
	price        float64
	brand        string
	discountRate int
}
