package function

func ToplaVariadic(sayilar ...int) int { // buradaki 3 nokta gelecek olan dedegiskenin variadik oldugunu soyluyor
	toplam := 0
	for i := 0; i < len(sayilar); i++ {
		toplam = toplam + sayilar[i]
	}
	return toplam
}
