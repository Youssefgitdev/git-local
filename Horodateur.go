package main

import "fmt"

func main() {
	var durerTempsRester float64
	var p10, p5, p2, p1, montantaPayer, argentDonné, monaie, rest int
	const prixHeure = 2
	//const x float64 = 1, 2, 5, 10

	fmt.Printf("Entrer le temps a passé : ")
	fmt.Scanln(&durerTempsRester)
	fmt.Printf("Montant a payer :")

	montantaPayer = int(durerTempsRester * prixHeure)
	if durerTempsRester < 0.5 {
		montantaPayer = 1
	}
	fmt.Println(montantaPayer)
	fmt.Printf("Argent Donné :")
	fmt.Scanln(&argentDonné)
	monaie = argentDonné - montantaPayer
	//if argentDonné != 10 || argentDonné != 5 || argentDonné != 2 || argentDonné != 1 {
	//fmt.Println("Erreur veuillez entrer une piece de 10 , 5 , 2")
	//}

	p10 = monaie / 10
	rest = monaie % 10
	p5 = rest / 5
	rest = rest % 5
	p2 = rest / 2
	p1 = rest / 1

	fmt.Println("votre monnaie :", monaie, "DH")
	fmt.Println("Nombre de piece de 10", p10, " de 5", p5, "de 2", p2, "de 1", p1)

}
