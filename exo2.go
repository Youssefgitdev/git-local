package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Entrer la valeur de a")
	fmt.Scanln(&a)
	fmt.Println("Entrer la valeur de b")
	fmt.Scanln(&b)
	fmt.Scanln(a * b)

	if a < 0 && b < 0 || a > 0 && b > 0 {
		fmt.Println(" resultat est positif")
	} else if a > 0 && b < 0 || a < 0 && b > 0 {
		fmt.Println(" resultat est negatif")
	} else if a == 0 || b == 0 {
		fmt.Println(" erreur")
	}
}
