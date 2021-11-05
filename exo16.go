package main

import (
	"fmt"
)

func main() {
	var n, max, pos int
	//var max = math.MaxInt
	max = 0
	for i := 0; i < 10; i++ {
		fmt.Println("Entrer les nombres")
		fmt.Scanln(&n)
		if i == 1 || n > max {
			max = n
			pos = i
		}
	}

	fmt.Println("la position de ", max, "est", pos)
}
