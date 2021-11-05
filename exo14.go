package main

import "fmt"

func main() {
	var x, n int
	fmt.Println("Entrer une valeur ")
	fmt.Scanln(&n)
	fmt.Scanln(x)
	x = 1
	for i := 2; i <= n; i++ {
		x = x * i
	}

	fmt.Println("La somme de 1 jusqu'a", n, "est", x)
}
