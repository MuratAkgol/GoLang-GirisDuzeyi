package slices

import "fmt"

func Demo1() {
	names := make([]string, 3)

	names[0] = "Murat"
	names[1] = "Cem"
	names[2] = "Şenay"
	names = append(names, "Akgöl")
	fmt.Println(names)
}
