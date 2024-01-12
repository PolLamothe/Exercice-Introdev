package decroissant

/*
La fonction decroissant doit trier un tableau d'entiers dans l'ordre décroissant.

# Entrée
- t, le tableau à trier (en place)

# Info
2022-2023, test 4, exercice 7
*/

func decalLeft(array *[]int, index int) {
	var result []int
	for i := 0; i < len(*array); i++ {
		if i == index {
			result = append(result, 0)
		}
		result = append(result, (*array)[i])
	}
	*array = result
}

func decroissant(tab *[]int) {
	var result []int = []int{}
	for i := 0; i < len(*tab); i++ {
		var lenResult int = len(result)
		for x := 0; x < lenResult; x++ {
			if (*tab)[i] >= result[x] {
				decalLeft(&result, x)
				result[x] = (*tab)[i]
				break
			} else if x == len(result)-1 {
				result = append(result, (*tab)[i])
			}
		}
		if len(result) == 0 {
			result = append(result, (*tab)[i])
		}
	}
	*tab = result
}
