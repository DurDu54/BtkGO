package error_handling

import (
	"fmt"
	"os"
)

func Demo1() {
	//dosya acaarken mainden konuma gitmak gerekiyor
	f, err := os.Open("demo1.txt")
	// dosya bulunursa error nill olarak doner

	if err != nil {
		//alttaki if type assetion
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("dosya bulunamadi : ", pErr.Path)
			return
		} else {
			fmt.Println("dosya bulunamadi")
			return
		}
	} else {
		fmt.Println("dosya bulundu :", f.Name())
	}

}
