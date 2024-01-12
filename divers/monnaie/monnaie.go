package monnaie

import (
	"errors"
	"math"
)

var errPasAssezPaye error = errors.New("le montant payé est inférieur au montant de l'achat, impossible de rendre la monnaie")

/*
Étant donnés un montant d'achat en euros et un montant payé en euros,
la fonction rendreMonnaie retourne la somme qu'il faut rendre au client.

# Entrées
- eurosAchat : la partie entière du montant d'achat en euros
- centimesAchat : la partie décimale du montant d'achat (donc les centimes)
- eurosPayes : la partie entière du montant payé en euros
- centimesPayes : la partie décimale du montant payé (donc les centimes)

# Sorties
- eurosRendus : la partie entière du montant à rendre en euros
- centimesRendus : la partie décimale du montant à rendre (donc les centimes)
- err : une erreur qui doit être errPasAssezPaye si le montant payé est inférieur au montant de l'achat et nil sinon

# Exemple
rendreMonnaie(12, 20, 15, 0) = 2, 80
*/
func rendreMonnaie(eurosAchat, centimesAchat, eurosPayes, centimesPayes int) (eurosRendus, centimesRendus int, err error) {
	var result float64 = (float64(eurosPayes) + float64(centimesPayes)*0.01) - (float64(eurosAchat) + float64(centimesAchat)*0.01)
	result = math.Round(result*100) / 100
	if result < 0 {
		return -1, -1, errPasAssezPaye
	}
	return int(result*100-float64(int(result*100)%100)) / 100, int(float64(int(result*100) % 100)), nil
}
