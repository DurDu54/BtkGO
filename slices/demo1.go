package slices

import (
	"fmt"
)

func Demo1() {

	isimler := make([]string, 3)
	fmt.Println(isimler)

	isimler[0] = "Mucahit"
	isimler[1] = "Durdu"
	isimler[2] = "Sakarya"
	isimler = append(isimler, "Adapazari")
	fmt.Println(isimler)
	fmt.Println(len(isimler))
}
