package defer_statement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) save() {
	fmt.Println("urun kaydedildi : ", p.productName)
	defer Log()
}
func Demo3() {
	p := product{productName: "laptop", unitPrice: 1000}
	defer p.save()
	p = product{productName: "mouse", unitPrice: 1000}
	/*
		burda kaydedşlen urun olarak laptop gozukur cunku defer method calıstıgında methodu ve gedecek olan parametreyı veya verıyı kaydedıyor ve method bıtınce cagırıyor
	*/
	fmt.Println("islem basarili")
}
func Log() {
	fmt.Println("loglandi")
}
