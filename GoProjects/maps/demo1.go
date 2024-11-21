package maps

import "fmt"

func Demo1() {
	dictionary := make(map[string]string)
	dictionary["Table"] = "Masa"
	dictionary["Book"] = "Kitap"
	dictionary["Pencil"] = "Kalem"

	fmt.Println(dictionary["Book"])
	fmt.Println("Emelman sayısı: ", len(dictionary))
	fmt.Println("Sözlük:", dictionary)

	delete(dictionary, "Book")
	fmt.Println("Emelman sayısı: ", len(dictionary))
	fmt.Println("Sözlük:", dictionary)

	//iki değer dönebilir.
	value, exist := dictionary["Table"]
	fmt.Println(value)
	fmt.Println("Listede olma durumu:", exist)

	dictionary2 := map[string]string{"Glass": "Bardak", "Microphone": "Mikrofon"}
	fmt.Println(dictionary2)
}
