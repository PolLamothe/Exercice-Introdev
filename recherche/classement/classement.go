package classement

/*
Étant donné un tableau d'entiers, la fonction classer doit les ranger dans trois tableaux : ceux qui sont pairs, ceux qui sont impairs et plus petits que 100, ceux qui sont impairs et plus grands que 100.

# Entrée
- t : un tableau d'entiers

# Sorties
- pairs : tous les entiers pairs de t
- petits : tous les entiers impairs plus petits que 100 de t
- grands : tous les entiers impairs plus grands que 100 de t

# Exemple
classer([]int{99, 100, 101}) = [100], [99], [101]
*/
func classer(t []int) (pairs, petits, grands []int) {
	var pairsTab []int
	var petitsTab []int
	var grandsTab []int
	for i := 0; i < len(t); i++ {
		if t[i]%2 == 0 {
			pairsTab = append(pairsTab, t[i])
		} else {
			if t[i] < 100 {
				petitsTab = append(petitsTab, t[i])
			} else {
				grandsTab = append(grandsTab, t[i])
			}
		}
	}
	return pairsTab, petitsTab, grandsTab
}
