package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, delta float64
	var x1 = -b + (math.Sqrt(delta))
	var x2 = -b - (math.Sqrt(delta))
	fmt.Println("ecrire la valeur de a ")
	fmt.Scanln(&a)
	fmt.Println("ecrire la valeur de b")
	fmt.Scanln(&b)
	fmt.Println("ecrire la valeur de c")
	fmt.Scanln(&c)
	delta = b*b - 4*a*c
	if delta > 0 {
		fmt.Println("y a 2 solutions")
		fmt.Println(x1 / 2 * a)
		fmt.Println(x2 / 2 * a)
	} else if delta == 0 {
		fmt.Println("y a une solution")
		fmt.Println(-b / (2 * a))
	} else {
		fmt.Println("ya pas de solution")
	}
}
