package valabs

import (
	"math"
)

/*
La fonction valabs doit trier un tableau d'entiers de la plus petite valeur absolue
à la plus grande valeur absolue. En cas
d'égalité de valeur absolue, les nombres
négatifs doivent être placés avant les
nombres positifs.

# Entrée
- tab : un tableau d'entiers
*/

func decalage(array []int, index int) []int {
	var result []int = []int{}
	for i := 0; i < len(array); i++ {
		if i == index {
			result = append(result, 0)
			result = append(result, array[i])
		} else {
			result = append(result, array[i])
		}
	}
	return result
}

func valabs(tab []int) []int {
	var result []int = []int{}
	for i := 0; i < len(tab); i++ {
		tempTab := math.Abs(float64((tab)[i]))
		if len(result) == 0 {
			result = append(result, (tab)[i])
		}
		var len int = len(result)
		if i != 0 {
			for x := 0; x < len; x++ {
				tempX := math.Abs(float64(result[x]))
				if tempX > tempTab {
					result = decalage(result, x)
					result[x] = (tab)[i]
					x = len
				}
				if tempX == tempTab {
					if result[x] >= (tab)[i] {
						result = decalage(result, x)
						result[x] = (tab)[i]
						x = len
					}
				} else if x == len-1 {
					result = append(result, (tab)[i])
				}
			}
		}
	}
	return result
}
