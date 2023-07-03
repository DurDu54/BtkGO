package main

import (
	"fmt"
	"golesson/arrays"
	"golesson/conditionals"
	"golesson/for_range"
	"golesson/function"
	"golesson/loops"
	"golesson/maps"
	"golesson/slices"
	"golesson/variables"
)

func main() {
	/*Lesson1()
	fmt.Println("****************************************************************")
	Lesson2()
	fmt.Println("****************************************************************")
	Lesson3()
	fmt.Println("****************************************************************")
	Lesson4()
	fmt.Println("****************************************************************")
	Lesson5()
	Lesson6()
	Lesson7()
	Lesson8()
	Lesson9()*/
	Lesson10()
}
func Lesson1() {
	// setup and Hello World
	fmt.Println("Hello World")
}
func Lesson2() {
	// variable

	// string
	var metin string = "Hello GO"
	fmt.Println(metin)
	fmt.Println(metin)
	fmt.Println(metin)
	fmt.Println(metin)

	// int
	var num int = 10
	fmt.Println(num)
	fmt.Println(num * 4)

	// float
	var num2 float32 = 0.1
	fmt.Println(num2)

	// tip belirsiz tanimlama
	num3 := 25
	fmt.Printf("Variable Type: %T\n", num3)

	// boolean
	var durum bool = false

	//if
	var metin1 string = "Engin"
	var metin2 string = "Mehmet"
	durum = metin1 == metin2
	fmt.Println(durum)
}
func Lesson3() {
	// bu ders paket ve modul olu;turma gosterildi
	// go mod init package_name yazarak module olusturulabilir
	// daha sonra impoort ederek cagrilip kullanilabilir
	variables.Demo1()
}
func Lesson4() {
	// bu ders if else anlatildi
	conditionals.Demo1()
}
func Lesson5() {
	// bu ders if else anlatildi
	loops.Demo1()
}
func Lesson6() {
	//arrays.Demo1()
	//arrays.Demo2()
	//arrays.Demo3()
	arrays.Demo4()
}
func Lesson7() {
	//slices.Demo1()
	slices.Demo2()
}
func Lesson8() {
	// function.SelamVer()
	// var asd = function.Topla(2, 3)
	// fmt.Println(asd)

	// sonuc1, sonuc2, sonuc3, sonuc4 := function.DortIslem(14, 3)
	// fmt.Printf("Toplam %v, Fark %v,Carpma %v, Bolme %v ", sonuc1, sonuc2, sonuc3, float32(sonuc4))

	//sonuclardan birini lamk icin sadece bir alt cizgi koymak yeterli oluyor
	//sonuc1, sonuc2, sonuc3, _ := gibi yapinca sonuc 4 gelmez

	sonuc1 := function.ToplaVariadic(3, 4, 5, 0)
	fmt.Println("sonuc :", sonuc1)

	sonuc2 := function.ToplaVariadic(3, 4, 0)
	fmt.Println("sonuc :", sonuc2)

	sonuc3 := function.ToplaVariadic(3, 5, 0)
	fmt.Println("sonuc :", sonuc3)

	sayilarArray := []int{6, 7, 8, 9}
	fmt.Println(function.ToplaVariadic(sayilarArray...)) // bu satirdaki uc nokta gonderilen dizinin variadic ildugunu soylemekte
}
func Lesson9() {
	maps.Demo1()
}
func Lesson10() {
	for_range.Demo1()
}
