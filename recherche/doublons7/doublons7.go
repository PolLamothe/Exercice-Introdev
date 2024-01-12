package doublons7

/*
On dispose d'un tableau d'entiers de longueur n et on suppose qu'il contient exactement les nombres de k à k + n - 1 (pas forcément dans l'ordre) pour un certain k non connu à l'avance. On voudrait vérifier que c'est bien le cas. C'est le travail de la fonction doublons.

# Entrée
- tab : un tableau d'entiers

# Sortie
- ok : un booléen qui doit valoir true si tab contient exactement les entiers de k à k + len(tab) - 1 et qui vaut false sinon

# Info
2023-2024, test 2, exercice 7
*/

func Trie(tab []int) []int {
	for i := 1; i < len(tab); i++ {
		if tab[i-1] > tab[i] {
			var temp int = tab[i]
			tab[i] = tab[i-1]
			tab[i-1] = temp
			return Trie(tab)
		}
	}
	return tab
}

func doublons(tab []int) (ok bool) {
	tab = Trie(tab)
	var k int = tab[0]
	if k+len(tab)-1 != tab[len(tab)-1] {
		return false
	}
	for i := 1; i < len(tab); i++ {
		if tab[i] != tab[i-1]+1 {
			return false
		}
	}
	return true
}
