package souschaine

import "strconv"

/*
Une chaine sc est une sous chaine d'une chaine s si on peut reconstruire s en rajoutant des lettres à sc. On peut le voir de manière récursive :
sc = asc' (avec a en caractère et sc' une chaine) est une sous chaine de s si s = as' et sc' est une sous chaine de s' ou si s = bs' et asc'
est une sous chaine de s'.

# Entrées
- s : une chaîne
- sc : une chaîne

# Sortie
- b : un booléen indiquant si sc est une sous chaine de s

# Exemple
sousChaine("abcde", "ace") = true
*/
func sousChaine(s, sc string) (b bool) {
	var allSumSC [][]string
	for i := 0; i < len(sc); i++ {
		var state bool = false
		for x := 0; x < len(allSumSC); x++ {
			if allSumSC[x][0] == string(sc[i]) {
				conversion, _ := strconv.Atoi(allSumSC[x][1])
				allSumSC[x][1] = strconv.Itoa(conversion + 1)
				state = true
			}
		}
		if !state {
			allSumSC = append(allSumSC, []string{string(sc[i]), "1"})
		}
	}
	print(allSumSC)
	for i := 0; i < len(allSumSC); i++ {
		var counter int
		for x := 0; x < len(s); x++ {
			if string(s[x]) == allSumSC[i][0] {
				counter++
			}
		}
		if strconv.Itoa(counter) != allSumSC[i][1] {
			return false
		}
	}
	return true
}
