package main

import "fmt"

func main() {
	var n, max, pos, i int
	//var max = math.MaxInt

	pos = 0
	n = 1
	max = 0
	for n > 0 {
		fmt.Println("Entrer les nombres")
		fmt.Scanln(&n)
		pos = pos + 1
		if pos == 1 || n > max {
			max = n
			i = pos
		} else if n == 0 {
			break
		}
	}

	fmt.Println("le nombre le plus grand est :", max)
	fmt.Println("Il a ete saisi en position :", i)
}
