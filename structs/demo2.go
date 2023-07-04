package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) save() {
	fmt.Println("eklendi :", c.firstName)

}
func (c customer) update() {
	fmt.Println("Guncellendi :", c.firstName)

}

func Demo2() {
	c := customer{firstName: "Mucahit", lastName: "Durdu", age: 22}
	c.save()
	c.update()
}
