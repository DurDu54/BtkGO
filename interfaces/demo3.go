package interfaces

import "fmt"

func Dogrula(i interface{}) {
	sayi, ok := i.(int)

	fmt.Println(sayi, ok)
}
func Demo3() {
	var sayi1 interface{} = "Mucahit"
	Dogrula(sayi1)
	var sayi2 interface{} = 4
	Dogrula(sayi2)
	var sayi3 interface{} = 7.16
	Dogrula(sayi3)

}
