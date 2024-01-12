package sousensembles

import (
	"errors"
	"fmt"
)

var errPasEnsemble error = errors.New("ceci n'est pas un ensemble")

/*
Étant donné un ensemble E on peut construire tous ses sous-ensembles de manière récursive en remarquant que si E' est E\{x} pour un élément x de E, alors les sous-ensembles de E sont les sous-ensembles de E' auxquels s'ajoutent ces mêmes sous-ensembles augmentés de x.
La fonction sousEnsembles doit mettre en œuvre cette construction.

# Entrée
- E : un tableau représentant un ensemble d'entiers

# Sortie
- PE : un tableau de tableaux contenant tous les sous-ensembles de E (sans doublons)
- err : une erreur qui vaut errPasEnsemble si et seulement si E n'est pas un ensemble (il contient plusieurs fois la même valeur ou il vaut nil)

# Exemple
sousEnsembles([]int{1, 2}) = [[] [1] [2] [1 2]] (l'ordre des ensembles et les ordres des valeurs dans les ensembles n'ont pas d'importance)
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

func isSame(big, small []int) bool {
	for i := 0; i < len(big); i++ {
		if big[i] != small[i] {
			return false
		}
	}
	return true
}

func getIndex(tab []int, value int) int {
	for i := 0; i < len(tab); i++ {
		if tab[i] == value {
			return i
		}
	}
	return -1
}

func recur(E []int, current [][]int, size int) (PE [][]int, err error) {
	fmt.Println(current, size)
	if len(current[len(current)-1]) == 0 {
		current = append(current, []int{E[0]})
		return recur(E, current, size)
	}
	var array []int
	if len(current[len(current)-1]) < 1 {
		array = []int{}
	} else {
		array = current[len(current)-1][1:]
	}
	if isSame(array, E[len(E)-len(current[len(current)-1]):]) {
		if size == len(E) {
			return current, nil
		}
		return recur(E, append(current, []int{E[0]}), size+1)
	} else {
		if current[len(current)-1][len(current[len(current)-1])-1] == E[len(E)-1] {
			current = append(current, []int{E[getIndex(E, current[len(current)-1][0])+1]})
			return recur(E, current, size)
		} else {
			current[len(current)-1] = append(current[len(current)-1], E[getIndex(E, current[len(current)-1][len(current[len(current)-1])-1])+1])
			return recur(E, current, size)
		}
	}
}

func sousEnsembles(E []int) (PE [][]int, err error) {
	if E == nil {
		return nil, errPasEnsemble
	}
	if len(E) == 0 {
		return [][]int{{}}, nil
	}
	for i := 0; i < len(E); i++ {
		if getOcc(E, E[i]) > 1 {
			return nil, errPasEnsemble
		}
	}
	return recur(E, [][]int{{}}, 1)
}
