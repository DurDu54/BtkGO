package arrays

import "fmt"

func Demo4() {
	var twoDAarray [2][3]int
	twoDAarray[0][0] = 1
	twoDAarray[0][1] = 3
	twoDAarray[0][2] = 5
	twoDAarray[1][0] = 2
	twoDAarray[1][1] = 4
	twoDAarray[1][2] = 6
	fmt.Println(twoDAarray)

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%v , ", twoDAarray[i][j])
		}
		fmt.Println()
	}
}
