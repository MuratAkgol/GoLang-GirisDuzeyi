package arrays

import "fmt"

func Demo2() {
	var cities [5]string
	cities[0] = "Ankara"
	cities[1] = "İstanbul"
	cities[2] = "İzmir"
	cities[3] = "Antalya"
	cities[4] = "Muğla"

	fmt.Println(cities)

	for i := 0; i < 5; i++ {
		fmt.Println(cities[i])
	}

}
