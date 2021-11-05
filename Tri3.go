package main

import "fmt"

func main() {
	var x, y, z, t int
	fmt.Printf("Donner le premier nombre : ")
	fmt.Scanln(&x)
	fmt.Printf("Donner le deuxième nombre : ")
	fmt.Scanln(&y)
	fmt.Printf("Donner le troisième nombre : ")
	fmt.Scanln(&z)
	if y < x {
		t = x
		x = y
		y = t
	}
	if z < x {
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
