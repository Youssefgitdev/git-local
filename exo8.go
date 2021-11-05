package main

import "fmt"

func main() {
	var hS, pU float64

	fmt.Println("Entrer la valeur du prix unitaire d'une heure")
	fmt.Scanln(&pU)
	fmt.Println("Entrer le nombre d'heures supplementaire")
	fmt.Scanln(&hS)

	if hS < 40 {
		fmt.Println("0")
	} else if hS >= 41 && hS <= 45 {
		fmt.Println((10 * pU) / 100)

	} else if hS >= 46 && hS <= 50 {
		fmt.Println((25 * pU) / 100)
	} else {
		fmt.Println((50 * pU) / 100)
	}
}
