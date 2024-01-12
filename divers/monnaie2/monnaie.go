package monnaie2

import (
	"errors"
)

var errPasAssezPaye error = errors.New("le montant payé est inférieur au montant de l'achat, impossible de rendre la monnaie")

/*
Étant donnés un montant d'achat en euros et un montant payé en euros,
la fonction rendreMonnaie retourne une liste de pièces et billets qui
permet de rendre au client la somme qu'il faut.

# Entrées
- eurosAchat : la partie entière du montant d'achat en euros
- centimesAchat : la partie décimale du montant d'achat (donc les centimes)
- eurosPayes : la partie entière du montant payé en euros
- centimesPayes : la partie décimale du montant payé (donc les centimes)

# Sorties
- eurosRendus : un tableau contenant les valeurs des pièces et billets en euros à rendre, ces valeurs doivent être prises parmi 1, 2, 5, 10, 20, 50, 100, 200 et 500
- centimesRendus : un tableau contenant les valeurs des pièces en centimes à rendre, ces valeurs doivent être prises parmi 1, 2, 5, 10, 20 et 50
- err : une erreur qui doit être errPasAssezPaye si le montant payé est inférieur au montant de l'achat et nil sinon

# Exemple
rendreMonnaie(12, 20, 15, 0) = [1, 1], [50, 20, 10] (ce n'est pas la seule possibilité pour ce rendu)
*/

func getSomme(tab []int) int {
	var sum int = 0
	for i := 0; i < len(tab); i++ {
		sum += tab[i]
	}
	return sum
}

func rendreMonnaie(eurosAchat, centimesAchat, eurosPayes, centimesPayes int) (eurosRendus, centimesRendus []int, err error) {
	var possibleEuro []int = []int{500, 200, 100, 50, 20, 10, 5, 2, 1}
	var possibleCentime []int = []int{50, 20, 10, 5, 2, 1}
	var euro []int = []int{}
	var centimes []int = []int{}
	if eurosAchat > eurosPayes {
		return nil, nil, errPasAssezPaye
	}
	if eurosAchat == eurosPayes && centimesAchat > centimesPayes {
		return nil, nil, errPasAssezPaye
	}
	var Achat float64 = float64(eurosAchat) + float64(centimesAchat)*0.01
	var Payé float64 = float64(eurosPayes) + float64(centimesPayes)*0.01
	var loop bool = true
	for float64(getSomme(euro))+Achat < Payé && loop {
		var state bool = false
		for i := 0; i < len(possibleEuro) && !state; i++ {
			if float64(getSomme(euro))+Achat+float64(possibleEuro[i]) <= Payé {
				euro = append(euro, possibleEuro[i])
				state = true
			}
		}
		if !state {
			loop = false
		}
	}
	for float64(getSomme(euro))+Achat+float64(getSomme(centimes))*0.01 < Payé {
		var state bool = false
		for i := 0; i < len(possibleCentime) && !state; i++ {
			if float64(getSomme(euro))+Achat+float64(getSomme(centimes))*0.01+(float64(possibleCentime[i])*0.01) <= Payé {
				centimes = append(centimes, possibleCentime[i])
				state = true
			}
		}
	}
	return euro, centimes, nil
}
