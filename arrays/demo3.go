package arrays

import "fmt"

func Demo3() {
	sayilar := [5]int{1, 2, 3, 4, 5}
	fmt.Println(sayilar)

	for i := 0; i < len(sayilar); i++ {
		fmt.Println(sayilar[i])
	}
}
