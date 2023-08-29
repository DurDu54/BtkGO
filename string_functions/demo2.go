package string_functions

import (
	p "fmt"
	s "strings"
)

func Demo2() {
	isim := "Mucahit"
	p.Println(s.HasPrefix(isim, "Muc"))
	p.Println(s.HasSuffix(isim, "hit"))

	harfler := []string{"M", "u", "c", "a", "h", "i", "t"}
	sonuc := s.Join(harfler, "*") //gelen stringleri yandaki tirnak icindekileri koyarak birlestirir
	p.Println(sonuc)

	p.Println(s.Replace(sonuc, "*", "+", len(sonuc)))
	p.Println(s.Split(sonuc, "-"))
	p.Println(s.Repeat(sonuc, 5))
}
