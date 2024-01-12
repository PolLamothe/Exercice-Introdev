package sousensembles2

import (
	"errors"
	"fmt"
)

var errPasEnsemble error = errors.New("ceci n'est pas un ensemble")

/*
Étant donné un ensemble E on peut construire tous ses sous-ensembles de manière récursive en remarquant que si E' est E\{x} pour un élément x de E, alors les sous-ensembles de E sont les sous-ensembles de E' auxquels s'ajoutent ces mêmes sous-ensembles augmentés de x.
La fonction sousEnsembles doit mettre en œuvre cette construction pour construire uniquement des sous-ensembles d'une longueur donnée.

# Entrée
- E : un tableau représentant un ensemble d'entiers
- k : un entier indiquant la taille des sous-ensembles à construire

# Sortie
- PE : un tableau de tableaux contenant tous les sous-ensembles de taille k de E (sans doublons)
- err : une erreur qui vaut errPasEnsemble si et seulement si E n'est pas un ensemble (il contient plusieurs fois la même valeur ou il vaut nil)

# Exemple
sousEnsembles([]int{1, 2}, 1) = [[1] [2]] (l'ordre des ensembles et les ordres des valeurs dans les ensembles n'ont pas d'importance)
*/

func getOcc(tab []int, value int) int {
	var count int
	for i := 0; i < len(tab); i++ {
		if tab[i] == value {
			count++
		}
	}
	return count
}

func recur(E []int, k int, current [][]int, ban []int) (PE [][]int, err error) {
	if len(current) == 0 && len(E) >= k {
		current = append(current, []int{})
		for i := 0; i < len(E) && len(current[0]) < k; i++ {
			current[0] = append(current[0], E[i])
		}
		return recur(E, k, current, []int{})
	} else if k > 0 && len(E) >= k {
		for i := 0; i < len(current); i++ {
			fmt.Println(ban)
			if current[i][0] == E[len(ban)] && current[i][len(current[i])-1] == E[len(ban)+k-1] {
				ban = append(ban, current[i][0])
			}
		}
		var copy []int
		for i := 0; i < len(E) && len(copy) < k; i++ {
			if getOcc(ban, E[i]) == 0 {
				copy = append(copy, E[i])
			}
		}
		if len(copy) < k {
			return current, nil
		} else {
			current = append(current, copy)
			return recur(E, k, current, ban)
		}
	}
	return current, nil
}

func sousEnsembles(E []int, k int) (PE [][]int, err error) {
	if E == nil {
		return nil, errPasEnsemble
	}
	for i := 0; i < len(E); i++ {
		if getOcc(E, E[i]) > 1 {
			return nil, errPasEnsemble
		}
	}
	return recur(E, k, [][]int{}, []int{})
}
