package main

import (
	"fmt"
	"golesson/arrays"
	"golesson/channels"
	"golesson/conditionals"
	"golesson/defer_statement"
	"golesson/error_handling"
	"golesson/for_range"
	"golesson/function"
	"golesson/interfaces"
	"golesson/loops"
	"golesson/maps"
	"golesson/pointers"
	"golesson/restfull"
	"golesson/slices"
	"golesson/string_functions"
	"golesson/structs"
	"golesson/variables"
)

func main() {
	/*
		Lesson1()
		Lesson2()
		Lesson3()
		Lesson4()
		Lesson5()
		Lesson6()
		Lesson7()
		Lesson8()
		Lesson9()
		Lesson10()
		Lesson11()
		Lesson12()
		Lesson13()
		Lesson14()
		Lesson15()
		Lesson16()
		Lesson17()

	*/
	Lesson18()
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
func Lesson11() {
	sayi := 20
	pointers.Demo1(&sayi)
	fmt.Println("maindeki sayi :", sayi)
	fmt.Println("************************************************")

	sayilar := []int{1, 2, 3}
	pointers.Demo2(sayilar)
	fmt.Println("maindeki sayilar :", sayilar[0])
}
func Lesson12() {
	//structs.Demo1()

	structs.Demo2()
}
func Lesson13() {
	//async calistirma basa go yazarak yapilir
	// go goroutines.CiftSayilar()
	// go goroutines.TekSayilar()
	// time.Sleep(5 * time.Second)

	// channel islemi bir async method calisinca geri deger almak icin kullaniliyor ve bu islemleri async olarak ilerliyor
	ciftSayiCn := make(chan int)
	tekSayiCn := make(chan int)
	go channels.CiftSayilar(ciftSayiCn)
	go channels.TekSayilar(tekSayiCn)

	ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn
	//  program channellarin bitmesini bekliyor ve ondan sonra devam ediyor channellar bitince devam ediyor
	//bu da bize async methodlardan sonra time.Sleep() kullanmadan methodun bitmesini bekliyor ve sonradan kulanilabiliyor
	carpim := ciftSayiToplam * tekSayiToplam
	fmt.Println(carpim)
}
func Lesson14() {
	//interfaces.Demo1()
	interfaces.Demo2()
}
func Lesson15() {
	defer_statement.Demo3()
}
func Lesson16() {
	error_handling.Demo3()
	//interfaces.Demo3()
}
func Lesson17() {
	// string_functions.Demo1()
	string_functions.Demo2()
}
func Lesson18() {
	//restfull.Demo1()
	restfull.Demo2()
}
