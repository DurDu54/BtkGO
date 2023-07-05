package defer_statement

import "fmt"

func A() {
	fmt.Println("a fonksiyonu calisti")
}

// defer ifadesi bir fonksiyondan calismasini istedigimiz fonksiyonnun ilk method bittikten sonra calismasini saglar
//fonksiyonnun neresine yazilirsa yazilsin her zaman en sonda calisir
// fonksiyon normalde yukardan assagi calisirfakat defer calisirken sondan baslar
// bunun sebebsi ise first in last out
/*
	b fonksiyonu calisti
	c fonksiyonu calisti
	a fonksiyonu calisti
*/
func B() {
	defer A()
	defer C()
	fmt.Println("b fonksiyonu calisti")

}
func C() {
	fmt.Println("c fonksiyonu calisti")
}
