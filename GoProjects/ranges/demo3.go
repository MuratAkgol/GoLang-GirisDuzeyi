package ranges

import "fmt"

func Demo3() {
	dictionary := map[string]string{"Book": "Kitap", "Table": "Masa"}

	for k, v := range dictionary {
		fmt.Println(k)
		fmt.Println(v)
	}
}
