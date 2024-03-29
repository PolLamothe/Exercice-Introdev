package doublons5

/*
On dispose d'un tableau d'entiers et on suppose qu'il ne contient pas de
doublons (aucun nombre n'est présent plusieurs fois dans le tableau.)
On voudrait vérifier cela, ce que va faire la fonction doublons.

# Entrée
- tab : un tableau d'entiers

# Sortie
- ok : un booléen qui vaut true si le tableau ne contient pas de doublons et
false sinon
*/

func doublons(tab []int) (ok bool) {
	for i := 0; i < len(tab); i++ {
		for x := 0; x < len(tab); x++ {
			if i != x && tab[i] == tab[x] {
				return false
			}
		}
	}
	return true
}
