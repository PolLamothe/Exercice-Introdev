package egalite

/*
On considère des multi-ensembles de nombres représentés par des tableaux :
ces tableaux peuvent contenir plusieurs fois chaque nombre (puisqu'ils
représentent des multi-ensembles) et les nombres ne sont pas nécessairement dans
l'ordre dans les tableaux.

On veut savoir si deux multi-ensembles sont égaux ou pas, c'est-a-dire savoir si
deux tableaux contiennent les mêmes nombres (et le même nombre de fois) ou pas.
C'est à cette question que doit répondre la fonction egalite.

# Entrées
- t1 : un tableau d'entiers représentant un multi-ensemble
- t2 : un tableau d'entiers représentant un multi-ensemble

# Sortie
- egaux : un booléen qui vaut true si t1 et t2 représentent le même multi-ensemble
          et qui vaut false sinon
*/

func egalite(t1, t2 []int) (egaux bool) {
	if (len(t1) == 0 || len(t2) == 0) && len(t1) != len(t2) {
		return false
	}
	for i := 0; i < len(t1); i++ {
		var state bool = false
		for x := 0; x < len(t2); x++ {
			if t1[i] == t2[x] {
				state = true
			}
		}
		if !state {
			return false
		}
	}
	return true
}
