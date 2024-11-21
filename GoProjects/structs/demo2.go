package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) saveCustomer() {
	fmt.Println("Kaydedildi.", c.firstName, c.lastName, c.age)
}
func (c customer) updateCustomer() {
	fmt.Println("Güncellendi.", c.firstName, c.lastName, c.age)
}

func Demo2() {
	c := customer{firstName: "Murat", lastName: "Akgöl", age: 24}

	c.updateCustomer()
}
