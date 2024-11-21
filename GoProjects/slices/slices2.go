package slices

import "fmt"

func Demo2() {
	cities := []string{"Ankara", "İstanbul", "İzmir"}
	fmt.Println(cities)

	copyOfCities := make([]string, len(cities))
	fmt.Println(copyOfCities)

	copy(copyOfCities, cities)
	fmt.Println(copyOfCities)

	//1. indexten 3. indekse kadar al.
	fmt.Println(cities[1:3])

	//birinci indexten sonuna kadar al
	fmt.Println(cities[1:])

	//en baştan ikinci indexi al
	fmt.Println(cities[:2])
}
