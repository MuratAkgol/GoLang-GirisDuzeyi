package pointers

import "fmt"

func Demo1(numb *int) {
	*numb++
	fmt.Println(*numb)
}
