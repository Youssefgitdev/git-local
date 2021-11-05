package main

import "fmt"

func main() {
	var s, n, i int

	fmt.Println("Entrer les nombre ")
	fmt.Scanln(&n)

	for i = 0; i < n; i++ {

		s = s + i*i
	}

	fmt.Println("La somme des carre naturels est :", s)
}
