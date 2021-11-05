package main

import (
	"fmt"
)

func main() {
	var n, max int
	//var max = math.MaxInt
	max = 0
	for i := 0; i < 10; i++ {
		fmt.Println("Entrer les nombres")
		fmt.Scanln(&n)
		if i == 1 || n > max {
			max = n
		}
	}

	fmt.Println("le max est :", max)
}
