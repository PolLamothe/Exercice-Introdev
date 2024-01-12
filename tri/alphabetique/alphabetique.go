package alphabetique

/*
La fonction alphabetique trie un tableau de chaînes de caractères dans l'ordre
alphabétique.

# Entrée
- dico : le tableau de chaînes de caractères à trier
*/

func alphabetique(dico []string) {
	for i := 1; i < len(dico); i++ {
		if dico[i] < dico[i-1] {
			dico[i], dico[i-1] = dico[i-1], dico[i]
			alphabetique(dico)
		}
	}
}
