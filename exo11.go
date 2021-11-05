package main

import "fmt"

func main() {
	var j int
	var m int
	var a int

	fmt.Println("Entrer le jour")
	fmt.Scanln(&j)
	fmt.Println("Entrer le mois")
	fmt.Scanln(&m)
	fmt.Println("Entrer l'année")
	fmt.Scanln(&a)

	switch {
	case j == 31 && m < 12:
		j = 1
		m++
		fmt.Println(j, m, a)
	case m == 2 && j == 28:
		j = 1
		m++
		fmt.Println(j, m, a)
	case m == 12 && j == 31:
		j = 1
		m = 1
		fmt.Println(j, m, (a + 1))
	case j > 31:

		j = 1
		fmt.Println(j, (m + 1), a)
	case m < 12:
		fmt.Println((j + 1), m, a)

	}

	//fmt.Println(j, m, a)
}
