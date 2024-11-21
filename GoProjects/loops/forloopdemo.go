package loops

import "fmt"

func Demo1() {
	for i := 0; i < 5; i++ {
		fmt.Println("Merhaba dünya " + fmt.Sprintf("%v", i))
	}
}
