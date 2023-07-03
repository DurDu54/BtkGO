package slices

import (
	"fmt"
)

func Demo2() {
	sehirler := []string{"ankara", "istanbul", "Sakarya"}
	fmt.Println(sehirler)
	sehirlerKopya := make([]string, len(sehirler))
	fmt.Println(sehirlerKopya)
	copy(sehirlerKopya, sehirler)
	fmt.Println(sehirlerKopya)
	//slice kopyalama islemi
	sehirler = append(sehirler, "adana")
	fmt.Println(sehirler)
	fmt.Println(sehirlerKopya)

	//slice parca parca gosterim ornek olarak apagination
	fmt.Println(sehirler[1:3])
	fmt.Println(sehirler[1:3])
	fmt.Println(sehirler[:2])
}
