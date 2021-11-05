package main

import "fmt"

func main() {
	var h int
	var m int

	fmt.Println("Entrer l'heure ")
	fmt.Scanln(&h)
	fmt.Println("Entrer la minutes")
	fmt.Scanln(&m)

	if h == 23 && m == 59 {
		h = 0
		m = 0
		fmt.Println(h, "heures", m, "minutes")
	} else if m == 59 {
		m = 0
		h++
		fmt.Println(h, "heures", m, "minutes")
	} else {
		h++
		m = 0
		fmt.Println(h, "heures", m, "minutes")
	}

}
