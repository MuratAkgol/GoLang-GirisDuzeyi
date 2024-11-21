package arrays

import "fmt"

func Demo4() {
	var multiArray [2][3]int

	multiArray[0][0] = 1
	multiArray[0][1] = 3
	multiArray[0][2] = 5
	multiArray[1][0] = 0
	multiArray[1][1] = 4
	multiArray[1][2] = 6

	fmt.Println(multiArray)
}
