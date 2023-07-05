package defer_statement

import "fmt"

/*
burda ciktida once bitti sonra Demo 2nin returnu yazar bunun sebebi ise
return oldugu icin geri donderiyor ve deferfuncu cagiri buda ciktinin boyle olmaina neden olur
*/
func Demo2(sayi int32) string {
	defer DeferFunc()

	if sayi%2 == 0 {
		return "cift sayidir"
	}

	if sayi%2 != 0 {
		return "cift sayidir"
	}
	return "belli degil"
}

func Test() {
	sonuc := Demo2(10)
	fmt.Println(sonuc)
}

func DeferFunc() {
	fmt.Println("Bitti")
}
