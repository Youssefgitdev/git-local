package main

import "fmt"

func main() {
	var x, n, r int
	fmt.Println("Entrer une valeur ")
	fmt.Scanln(&n)
	fmt.Scanln(r)

	for r := 1; r <= n; r++ {
		x = x + r
	}

	fmt.Println("La somme de 1 jusqu'a", n, "est", x)
}
