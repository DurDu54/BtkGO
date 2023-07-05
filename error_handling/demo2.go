package error_handling

import (
	"errors"
	"fmt"
)

func TahminEt(tahmin int) (string, error) {
	aklimdakiSayi := 50
	if tahmin < 1 || tahmin >= 100 {
		return "", errors.New("1-100 arasinda bir sayi giriniz.")
	}
	if tahmin > aklimdakiSayi {
		return "", errors.New("daha kucuk bir sayi giriniz")
	}
	if tahmin < aklimdakiSayi {
		return "", errors.New("daha buyuk bir sayi giriniz")
	}
	return " bildiniz", nil
}

func Demo2() {
	mesaj, hata := TahminEt(53)
	fmt.Println(mesaj)
	fmt.Println(hata)

}
