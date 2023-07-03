package for_range

import "fmt"

func Demo1() {
	sehirler := []string{"ankara", "istanbul", "izmir", "Sakarya"}
	for i := 0; i < len(sehirler); i++ {
		fmt.Println(sehirler[i])
	}
	fmt.Println("****************************************************************")
	for _, sehir := range sehirler { // _ yerine bir har fyazilarak hangi indiste oldugumuzun kontrolu yapilabilir yada hicbirsey yazilmadan direkt olarak gecilarbilir
		fmt.Println(sehir)
	}
	fmt.Println("****************************************************************")
	sayilar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	toplam := 0
	for _, value := range sayilar {
		if value%2 == 1 {
			toplam = toplam + value
		}
	}
	fmt.Println(toplam)
	fmt.Println("****************************************************************")
	sozluk := map[string]string{"book": "kitap", "table": "masa"}
	for k, v := range sozluk { // buradaki gibi bir sozluk yapisinda kullanilirsa da hem keyi hem valueyi verir
		fmt.Print(k)
		fmt.Println("  ", v)
	}
}
