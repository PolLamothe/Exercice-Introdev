package tri

/*
La fonction tri doit trier un tableau d'entiers de la manière suivante : on trouvera d'abord tous les nombres pairs du tableau, dans l'ordre croissant, puis tous les nombres impairs, dans l'ordre décroissant.

# Entrée
- t, le tableau à trier (en place)

# Info
2022-2023, test 2, exercice 4
*/

func triCroissant(tab []int) []int {
	for i := 1; i < len(tab); i++ {
		if tab[i] <= tab[i-1] {
			var temp int = tab[i-1]
			tab[i-1] = tab[i]
			tab[i] = temp
			if tab[i] != tab[i-1] {
				return triCroissant(tab)
			}
		}
	}
	return tab
}
func triDecroissant(tab []int) []int {
	for i := 1; i < len(tab); i++ {
		if tab[i] >= tab[i-1] {
			var temp int = tab[i-1]
			tab[i-1] = tab[i]
			tab[i] = temp
			if tab[i] != tab[i-1] {
				return triDecroissant(tab)
			}
		}
	}
	return tab
}

func tri(t []int) []int {
	var pairTab []int
	var impairTab []int
	var result []int
	for i := 0; i < len(t); i++ {
		if t[i]%2 == 0 {
			pairTab = append(pairTab, t[i])
			pairTab = triCroissant(pairTab)
		}
		if t[i]%2 == 1 {
			impairTab = append(impairTab, t[i])
			impairTab = triDecroissant(impairTab)
		}
	}
	for i := 0; i < len(pairTab); i++ {
		result = append(result, pairTab[i])
	}
	for i := 0; i < len(impairTab); i++ {
		result = append(result, impairTab[i])
	}
	return result
}
