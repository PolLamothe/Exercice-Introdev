package recherche

/*
La fonction recherche doit indiquer la valeur et la position du plus petit
entier strictement positif dans un tableau.

# Entrée
- tab : le tableau dans lequel on cherche

# Sorties
- trouve : un booléen qui vaut true s'il existe au moins un entier strictement
           positif dans tab et false sinon,
- pos : la position dans tab du plus petit entier strictement positif s'il existe
- val : la valeur du plus petit entier strictement positif de tab s'il existe
(pos et val auront la valeur par défaut d'un entier s'il n'existe pas d'entier
strictement positif dans tab)

# Info
2021-2022, test2, exercice 2
*/

func recherche(tab []int) (trouve bool, pos, val int) {
	var min int
	var posi int = -1
	for i := 0; i < len(tab); i++ {
		if posi == -1 && tab[i] > 0 {
			min = tab[i]
			posi = i
		} else if tab[i] > 0 {
			if tab[i] < min {
				min = tab[i]
				posi = i
			}
		}
	}
	if posi != -1 {
		return true, posi, min
	}
	return false, 0, 0
}
