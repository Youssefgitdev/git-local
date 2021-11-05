package main

import "fmt"

func main() {
	var x, y, z int

	fmt.Println("Ecrire la valeur de x")
	fmt.Scanln(&x)
	fmt.Println("Ecrire la valeur de y")
	fmt.Scanln(&y)
	fmt.Println("Ecrire la valeur de z")
	fmt.Scanln(&z)

	if x <= y && y <= z {
		fmt.Println(" x , y ,z sont croissant")
	} else {
		fmt.Println("x , y , z ne sont pas croissant")
	}
}
