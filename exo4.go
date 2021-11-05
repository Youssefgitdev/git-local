package main

import "fmt"

func main() {
	var x, y, z, t int
	fmt.Println("entrer la valeur de x ")
	fmt.Scanln(&x)
	fmt.Println("entrer la valeur de y")
	fmt.Scanln(&y)
	fmt.Println("entrer la valeur de z")
	fmt.Scanln(&z)

	if y < x {
		t = x
		x = y
		y = t
	} else if z < x {
		t = z
		z = y
		y = x
		x = t

	} else if z < y {
		t = z
		z = y
		y = t
	}
	fmt.Println(x, y, z)
}
