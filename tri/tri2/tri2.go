package tri2

import "fmt"

/*
La fonction triabs doit trier un tableau d'entiers de la plus grande valeure
absolue à la plus petite valeure absolue. Cette fonction ne doit pas modifier
le tableau donné en entrée.

# Entrée
- tinit : un tableau d'entiers qui ne doit pas être modifié.

# Sortie
- tfin : un tableau contenant les mêmes entiers que tinit mais triés du plus
         grand (en valeure absolue) au plus petit (en valeure absolue).

# Info
2021-2022, test2, exercice 7
*/

func decalLeft(array *[]int, index int) {
	var copy []int = []int{}
	for i := 0; i < len(*array); i++ {
		if i == index {
			copy = append(copy, 0)
		}
		copy = append(copy, (*array)[i])
	}
	*array = copy
}

func getAbsoluteValue(number int) int {
	if number < 0 {
		return (number * number) / -number
	} else {
		return number
	}
}

func triabs(tinit []int) []int {
	var tfin []int = []int{}
	for i := 0; i < len(tinit); i++ {
		var iAbsolute int = getAbsoluteValue(tinit[i])
		var tfinLen int = len(tfin)
		for x := 0; x < tfinLen; x++ {
			var xAbsolute int = getAbsoluteValue(tfin[x])
			if xAbsolute <= iAbsolute {
				decalLeft(&tfin, x)
				tfin[x] = tinit[i]
				break
			}
			if x == tfinLen-1 {
				tfin = append(tfin, tinit[i])
			}
		}
		if tfinLen == 0 {
			tfin = append(tfin, tinit[i])
		}
	}
	fmt.Println(tfin)
	return tfin
}
