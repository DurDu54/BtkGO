package main

import (
	"fmt"
	"golesson/conditionals"
	"golesson/variables"
)

func main() {
	Lesson1()
	fmt.Println("****************************************************************")
	Lesson2()
	fmt.Println("****************************************************************")
	Lesson3()
	fmt.Println("****************************************************************")
	Lesson4()
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
