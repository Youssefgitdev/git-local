package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("entrer la valeur de a")
	fmt.Scanln(&a)
	fmt.Println("entrer la valeur de b")
	fmt.Scanln(&b)

	if a != 0 {
		fmt.Println(-b / a)

	} else {
		fmt.Println("pas de solution")
	}
}
