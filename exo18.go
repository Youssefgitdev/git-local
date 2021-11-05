package main

import "fmt"

func main() {
	var n, max, s, min, total int
	//var max = math.MaxInt
	total = 0

	s = 0

	fmt.Println("Entrer les nombres")
	fmt.Scanln(&n)
	s = n
	min = n
	max = n
	total = total + 1
	for n != 0 {
		fmt.Println("Entrer les nombres")
		fmt.Scanln(&n)
		total = total + 1

		s = s + n
		if n > max {
			max = n
			if n < min {
				min = n
			}
		}

	}

	//total = total - 1
	fmt.Println("le nombre le plus grand est :", max)
	fmt.Println("le nombre le plus petit est :", min)
	fmt.Println("la somme est :", s)
	fmt.Println("le total des nombres est :", total-1)
	fmt.Println("La moyenne est :", s/(total-1))

}
