package souschaine

import (
	"fmt"
	"strconv"
)

/*
Une chaine sc est une sous chaine d'une chaine s si on peut reconstruire s en rajoutant des lettres à sc. Une chaine sc est une sous chaine commune à s1 et s2 si c'est une sous chaine de s1 et une sous chaine de s2. C'est une plus longue sous chaine commune si on ne peut pas trouver sc2 qui soit aussi une sous chaine commune à s1 et s2 et qui ait strictement plus de lettres que sc.

Un algorithme a été vu en cours pour calculer la longueur de plus longues sous chaines communes. L'objet de cet exercice est de l'adapter pour obtenir, en plus de leur longueur, les sous chaines elles-mêmes.

# Entrées
- s1 : une chaîne
- s2 : une chaîne

# Sortie
- s : une plus longue sous chaine commune à s1 et s2

# Exemple
sousChaine("bonjour", "bonsoir") = bonor
*/
func sousChaine(s1, s2 string) (s string) {
	var allSumSC [][]string
	for i := 0; i < len(s1); i++ {
		var state bool = false
		for x := 0; x < len(allSumSC); x++ {
			if allSumSC[x][0] == string(s1[i]) {
				conversion, _ := strconv.Atoi(allSumSC[x][1])
				allSumSC[x][1] = strconv.Itoa(conversion + 1)
				state = true
			}
		}
		if !state {
			allSumSC = append(allSumSC, []string{string(s1[i]), "1"})
		}
	}
	var result string = ""
	fmt.Println(allSumSC)
	for i := 0; i < len(s2); i++ {
		for x := 0; x < len(allSumSC); x++ {
			if string(s2[i]) == allSumSC[x][0] && allSumSC[x][1] != "0" {
				conversion, _ := strconv.Atoi(allSumSC[x][1])
				allSumSC[x][1] = strconv.Itoa(conversion - 1)
				result = result + string(s2[i])
			}
		}
	}
	return result
}
