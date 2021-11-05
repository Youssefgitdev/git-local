package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Println(" entrer la 1ere interval")
	fmt.Scanln(&a)
	fmt.Println("entrer la 2eme interval")
	fmt.Scanln(&b)
	fmt.Println(" entrer la valeur a verifier")
	fmt.Scanln(&c)

	if c >= a && c <= b {
		fmt.Println("c appartient a l'interval")
	} else {
		fmt.Println("c n'appartient pas a l'interval")
	}

}
