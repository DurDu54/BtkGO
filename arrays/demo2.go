package arrays

import "fmt"

func Demo2() {
	fmt.Println("diziye eleman girisi yapin")
	var dizi [5]int
	for i := 0; i < 5; i++ {
		fmt.Printf("%v. sayiyi giriniz: ", i+1)
		fmt.Scan(&dizi[i])
		fmt.Println()
	}
	for i := 0; i < 5; i++ {
		fmt.Println(dizi[i])
	}
}
