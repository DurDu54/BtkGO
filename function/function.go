package function

import "fmt"

func Topla(sayi1 int, sayi2 int) int {
	var toplam = sayi1 + sayi2
	fmt.Println("Sonuc: ", toplam)
	return toplam
}
func SelamVer() {
	fmt.Println("Merhaba")
}

func DortIslem(sayi1 int, sayi2 int) (int, int, int, float32) {
	Toplam := sayi1 + sayi2
	Fark := sayi1 - sayi2
	carpim := sayi1 * sayi2
	bolum := float32(sayi1 / sayi2)
	return Toplam, Fark, carpim, bolum
}
