package ranges

import "fmt"

func Demo1() {
	cities := []string{"Ankara", "İstanbul", "İzmir"}
	// for i := 0; i < len(cities); i++ {
	// 	fmt.Println(cities[i])
	// }

	//bu şekilde de for dönebilirsin
	for _, city := range cities {
		fmt.Println(city)
	}
}
