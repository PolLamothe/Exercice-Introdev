package doublons8

/*
On dispose d'un tableau d'entiers de longueur n et on suppose qu'il contient exactement p fois chaque nombre compris entre k et k + (n/p) - 1 (pas forcément dans l'ordre) pour un certain k et un certain p non connus à l'avance (p différent de 0). On voudrait vérifier que c'est bien le cas. C'est le travail de la fonction doublons.

# Entrée
- tab : un tableau d'entiers

# Sortie
- ok : un booléen qui doit valoir true si tab contient exactement p fois chaque entier de k à k + len(tab)/p - 1 et qui vaut false sinon

# Info
2023-2024, test 2, exercice 8
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
	var p int
	for i := 0; i < len(tab); i++ {
		if tab[i] == k {
			p++
		} else {
			break
		}
	}
	if p == 0 {
		return false
	}
	if tab[len(tab)-1] != k+len(tab)/p-1 {
		return false
	}
	for i := 0; i < len(tab); i++ {
		var counter int = 0
		for x := 0; x < len(tab); x++ {
			if tab[i] == tab[x] {
				counter++
			}
		}
		if counter != p {
			return false
		}
	}
	return true
}
