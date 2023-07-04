package goroutines

import (
	"fmt"
	"time"
)

func CiftSayilar() {
	for i := 0; i < 10; i += 2 {
		fmt.Println("Cift Sayi : ", i)
		time.Sleep(time.Second)
	}
}
func TekSayilar() {
	for i := 1; i < 10; i += 2 {
		fmt.Println("Tek Sayi : ", i)
		time.Sleep(time.Second)
	}
}
