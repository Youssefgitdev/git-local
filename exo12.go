package main

import "fmt"

func main() {
	var x, n int

	fmt.Println("Entrer un nombre")
	fmt.Scanln(&n)
	fmt.Println("Les nombres sont : ")
	for x = n + 1; x < n+11; x++ {
		fmt.Println(x)
	}

}
