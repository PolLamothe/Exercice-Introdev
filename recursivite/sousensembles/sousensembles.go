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

func recur(E, previous []int, length int) [][]int {
	if len(E) == 0 {
		return [][]int{}
	}
	var result [][]int = [][]int{}
	if length == 4 && len(previous) == 3 {
		fmt.Println(E, previous)
	}
	if len(previous) == length-1 {
		for i := 0; i < len(E); i++ {
			if length == 4 && len(previous) == 3 {
				fmt.Println(result)
			}
			temp2 := append([]int(nil), previous...)
			temp := append(temp2, E[i])
			if length == 4 && len(previous) == 3 {
				fmt.Println(result)
			}
			result = append(result, temp)
			if length == 4 && len(previous) == 3 {
				fmt.Println(result)
			}
		}
	} else {
		for i := 0; i < len(E); i++ {
			result2 := recur(E[i+1:], append(previous, E[i]), length)
			result = append(result, result2...)
		}
	}

	return result
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
	var result [][]int = [][]int{{}}
	for i := 1; i <= len(E); i++ {
		result = append(result, recur(E, []int{}, i)...)
	}
	fmt.Println(recur(E, []int{}, 4))
	return result, nil
}
