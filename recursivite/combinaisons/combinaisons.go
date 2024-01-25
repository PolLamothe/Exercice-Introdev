package combinaison

import "fmt"

/*
Étant donné un entier n, on souhaite trouver toutes les combinaisons d'entiers (c'est-à-dire des ensemble avec répétitions) compris entre 1 et n (inclus) dont la somme vaut n. Ainsi, par exemple, pour n = 3 on doit trouver les combinaisons {1, 1, 1}, {1, 2} et {3}.

Il est possible d'écrire pour calculer les combinaisons pour un entier n une fonction récursive sur un principe proche du suivant :
* Pour chaque i entre 1 et n
- calculer les combinaisons pour n-i
- ajouter à chacune d'entre elles la valeur i
* Retourner toutes les combinaisons obtenus

En le faisant, vous remarquerez que vous construisez trop de combinaisons (il y a des doublons). Pour éviter cela, vous pouvez utiliser une borne indiquant la valeur maximum à considérer (qui sera donc le minimum entre n et cette borne), qui sera passée dans vos appels récursifs.

Le plus simple est d'écrire une fonction récursive, mais les boucles for sont autorisées.

# Entrées
- n, un entier
- res, un tableau de tableaux d'entiers, chacun de ses tableaux représente une combinaison d'entiers compris entre 1 et n et dont la somme vaut n. L'ordre des combinaisons n'a pas d'importance.

# Info
2022-2023, test 2, exercice 8
*/

func isInArray(index int, array []int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == index {
			return true
		}
	}
	return false
}

func pop(array [][]int, index []int) [][]int {
	var result [][]int
	for i := 0; i < len(array); i++ {
		if !isInArray(i, index) {
			result = append(result, array[i])
		}
	}
	return result
}

func getSum(array []int) int {
	var sum int = 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum
}

func GetBiggest(array []int) int {
	var max int = 0
	for i := 0; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
	}
	return max
}

func sameContent(array1, array2 []int) bool {
	if len(array1) != len(array2) {
		return false
	}
	if GetBiggest(array1) != GetBiggest(array2) {
		return false
	}
	var result []int = make([]int, GetBiggest(array1))
	for i := 0; i < len(array1); i++ {
		result[array1[i]-1]++
	}
	for i := 0; i < len(array2); i++ {
		result[array2[i]-1]--
	}
	for i := 0; i < len(result); i++ {
		if result[i] != 0 {
			return false
		}
	}
	return true
}

func recur(previous []int, length, n int) [][]int {
	var copy []int = append([]int{}, previous...)
	for i := len(copy) - 1; i >= 0; i-- {
		if copy[i] < n {
			copy[i]++
			break
		} else {
			copy[i] = 1
		}
	}
	var state bool = true
	for i := 0; i < len(copy); i++ {
		if copy[i] != n {
			state = false
		}
	}
	if state {
		return [][]int{copy}
	}
	return append(recur(copy, length, n), copy)
}

func combinaisons(n int) (res [][]int) {
	var allCombinaisons [][]int
	for i := 1; i <= n; i++ {
		var previous []int = []int{}
		for x := 0; x < i; x++ {
			previous = append(previous, 1)
		}
		previous[len(previous)-1] = 0
		var copy []int = append([]int{}, previous...)
		var temp [][]int = recur(copy, i, n)
		allCombinaisons = append(allCombinaisons, temp...)
	}
	var banIndex []int = []int{}
	for i := 0; i < len(allCombinaisons); i++ {
		if getSum(allCombinaisons[i]) != n {
			banIndex = append(banIndex, i)
		}
	}
	allCombinaisons = pop(allCombinaisons, banIndex)
	banIndex = []int{}
	for j := 0; j < len(allCombinaisons); j++ {
		for x := 0; x < len(allCombinaisons); x++ {
			if j != x {
				if sameContent(allCombinaisons[j], allCombinaisons[x]) {
					if j < x {
						banIndex = append(banIndex, x)
					} else {
						banIndex = append(banIndex, j)
					}
				}
			}
		}
	}
	allCombinaisons = pop(allCombinaisons, banIndex)
	fmt.Println(allCombinaisons)
	return allCombinaisons
}
