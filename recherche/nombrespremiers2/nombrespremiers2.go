package nombrespremiers2

import (
	"fmt"
	"math"
)

/*
La fonction selectionPremiers filtre le contenu d'un tableau d'entiers pour ne garder que ceux qui sont des nombres premiers, en éliminant les doublons.

# Entrée
- t : un tableau d'entiers

# Sortie
- p : un tableau contenant tous les nombres premiers de t, sans doublons.
Si t est vide, p doit être identique à t.

# Exemple
selectionPremiers([]int{1, 2, 2, 3, 4, 5}) = [2 3 5] (l'ordre n'a pas d'importance)
*/

func isInArray(array []int, number int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == number {
			return true
		}
	}
	return false
}

func selectionPremiers(t []int) (p []int) {
	if len(t) == 0 {
		return t
	}
	var result []int = []int{}
	for i := 0; i < len(t); i++ {
		var state bool = false
		for x := t[i]; x > 0; x-- {
			if x != t[i] && x != 1 {
				if math.Mod(float64(t[i])/float64(x), 1) == 0 {
					state = true
				}
			}
		}
		if !state && t[i] > 1 && !isInArray(result, t[i]) {
			result = append(result, t[i])
		}
	}
	fmt.Println(result)
	return result
}
