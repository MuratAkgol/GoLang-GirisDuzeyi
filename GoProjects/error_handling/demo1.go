package error_handling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("demo1.txt")
	// Dosya bulunursa err nil değer atılır.

	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Dosya bulunamadı.", pErr.Path)
			return
		} else {
			fmt.Println("Dosya bulunamadı.", pErr.Path)
			return
		}
	} else {
		fmt.Println(f.Name())
	}

}
