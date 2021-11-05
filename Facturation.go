package main

import "fmt"

func main() {
	var PrixUnitaire, Quantite, TauxRemise, Remise float64
	var TotalAvantRemise, totalTTC, TotalHT, MontantTVA float64
	const TauxTVA float64 = 0.2

	fmt.Printf("Entrer le prix Unitaire : ")
	fmt.Scanln(&PrixUnitaire)
	fmt.Printf("Entrer la quantité :")
	fmt.Scanln(&Quantite)
	fmt.Printf("Entrer le taux de remise : ")
	fmt.Scanln(&TauxRemise)

	TotalAvantRemise = PrixUnitaire * Quantite
	Remise = TotalAvantRemise * TauxRemise
	TotalHT = TotalAvantRemise - Remise
	MontantTVA = TotalHT * TauxTVA
	totalTTC = TotalHT + MontantTVA
	fmt.Println("Total avant Remise :", TotalAvantRemise)
	fmt.Println("Montant de la remise:", Remise)
	fmt.Println("Total HT:", TotalHT)
	fmt.Println("Montant TVA :", MontantTVA)
	fmt.Println("Le montant TTC est :", totalTTC, "DH")
}
