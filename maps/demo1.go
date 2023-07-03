package maps

import "fmt"

func Demo1() {
	//map bir sozluk yapisidir
	// yada switch keys gibi birsey
	// enumada benziyor

	//map yapisinda key-value yapisi vardir
	sozluk := make(map[string] /*key*/ string /*value*/)
	sozluk["table"] = "masa"
	sozluk["book"] = "kitap"
	sozluk["title"] = "baslik"
	fmt.Println("sozlugumuz: ", sozluk)
	fmt.Println(sozluk["table"])
	fmt.Println(len(sozluk))
	delete(sozluk, "table")
	fmt.Println("sozlugumuz: ", sozluk)
	fmt.Println(len(sozluk))

	deger, varMi := sozluk["title"]
	fmt.Println("deger: ", deger)
	fmt.Println("varMi: ", varMi)
}
