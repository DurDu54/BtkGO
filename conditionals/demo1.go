package conditionals

import "fmt"

func Demo1() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen > hesap { //
		fmt.Println("Para Yetersiz")

	} else {
		fmt.Println("paraniz veriliyor")
	}
}
