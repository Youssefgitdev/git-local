package main

import "fmt"

func main() {
	var n int
	fmt.Println("Entrer la note")
	fmt.Scanln(&n)
	if n < 0 {
		fmt.Println(" erreur ")
	}
	switch {
	case n >= 0 && n < 10:
		fmt.Println("Echec")
	case n >= 10 && n < 12:
		fmt.Println("Mention passable")
	case n >= 12 && n < 14:
		fmt.Println("Mention assez bien")
	case n >= 14 && n < 16:
		fmt.Println("Mention bien")
	case n >= 16 && n <= 20:
		fmt.Println("Mention très bien")

	}

}
