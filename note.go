package main

import "fmt"

func main() {
	var n float64
	fmt.Println("Entrer la note :")
	fmt.Scanln(&n)
	if n < 0 {
		fmt.Println("err")
	}
	if n >= 0 && n < 10 {
		fmt.Println("Echec")
	} else if n < 12 {
		fmt.Println("Mention passable")
	} else if n < 14 {
		fmt.Println("Mention assez bien")
	} else if n < 16 {
		fmt.Println("Mention bien")
	} else if n <= 20 {
		fmt.Println("Mension Très bien")
	} else {
		fmt.Println("Erreur")
	}

}
