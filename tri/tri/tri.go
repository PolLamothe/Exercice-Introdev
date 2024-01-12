package tri

import "fmt"

/*
La fonction tri doit trier un tableau d'entiers du plus petit au plus grand.
Cette fonction ne doit pas modifier le tableau donné en entrée.

# Entrée
- tinit : un tableau d'entiers qui ne doit pas être modifié.

# Sortie
- tfin : un tableau contenant les mêmes entiers que tinit mais triés du plus
         petit au plus grand.

# Info
2021-2022, test2, exercice 6
*/

func decalLeft(tabs *[]int, index int) {
	var result []int
	for i := 0; i < len(*tabs); i++ {
		if i == index {
			result = append(result, 0)
		}
		result = append(result, (*tabs)[i])
	}
	*tabs = result
}

func tri(tinit []int) (tfin []int) {
	var result []int
	for i := 0; i < len(tinit); i++ {
		if len(result) == 0 {
			result = append(result, tinit[i])
		} else {
			var resultLen = len(result)
			for x := 0; x < resultLen; x++ {
				if result[x] > tinit[i] {
					decalLeft(&result, x)
					result[x] = tinit[i]
					x = resultLen
				}
				if x == (resultLen - 1) {
					result = append(result, tinit[i])
					break
				}
			}
		}
	}
	fmt.Println(result)
	return result
}
