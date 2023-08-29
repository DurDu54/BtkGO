package string_functions

import (
	"fmt"
	s "strings"
) // ustteki s "string" aslinda string. d'yerek string methodlarini getirirken artik sadece s yazdigimizda gelmesini sagliyor Buna ALIAS Deniyor

func Demo1() {
	isim := "Mucahit"
	fmt.Println(s.Count(isim, "u"))    // substr kac kere geciyor
	fmt.Println(s.Contains(isim, "u")) // substr var mi
	fmt.Println(s.Index(isim, "u"))    // substr gordugu ilk index (diziler gibi 0den basliyor)
	fmt.Println(s.Index(isim, "t"))    // eger yoksa -1 bir doner
	fmt.Println(s.ToLower(isim))
	fmt.Println(s.ToUpper(isim))
}
