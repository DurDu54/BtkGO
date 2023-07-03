package loops

import "fmt"

func Demo1() {
	sayi := 0
	toplam := 0
	for i := 0; i < 5; i++ {
		fmt.Printf("%v. sayiyi giriniz: ", i+1)
		fmt.Scanln(&sayi)
		toplam += sayi
	}
	fmt.Printf("Sayilariniz toplam %v yapiyor.", toplam)
}
