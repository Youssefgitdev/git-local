package main

import "fmt"

func main() {
	var h int
	var m int
	var s int

	fmt.Println("Entrer l'heure ")
	fmt.Scanln(&h)
	fmt.Println("Entrer la minutes")
	fmt.Scanln(&m)
	fmt.Println("Entrer une seconde")
	fmt.Scanln(&s)

	if h == 23 && m == 59 && s == 59 {
		h = 0
		m = 0
		s = 0
		fmt.Println(h, "heures", m, "minutes", s, "secondes")

	} else if h == 23 && m == 59 && s < 60 {
		s++
		fmt.Println(h, "heures", m, "minutes", s, "secondes")

	} else if h == 23 && m == 59 {
		h = 0
		m = 0
		fmt.Println(h, "heures", m, "minutes")
	} else if m == 59 && s == 59 {
		s = 0
		m = 0
		h++
		fmt.Println(h, "heures", m, "minutes", s, "secondes")

	} else if s == 59 {
		s = 0
		m++
		fmt.Println(h, "heures", m, "minutes", s, "secondes")

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
