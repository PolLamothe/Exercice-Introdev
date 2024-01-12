package piecesjaunes

import (
	"errors"
	"fmt"
)

var errImpossible error = errors.New("impossible de constituer cette somme avec ces pièces")

/*
Les pièces et billets que nous utilisons dans la vie de tous les jours ont des valeurs qui ont été choisies pour qu'il soit simple de
rendre la monnaie (il suffit de toujours donner la pièce ou le billet de plus grande valeur possible, jusqu'à avoir rendu la somme
voulue). Dans cet exercice, on imagine que ce n'est pas le cas : on dispose d'un ensemble de valeurs de pièces quelconque, on a accès
à autant de chacune de ces pièces qu'on le souhaite, et on cherche à constituer une somme avec le moins de pièces possible.

La fonction meilleurDecomposition doit trouver la meilleur décomposition d'une somme d'argent (celle qui utilise le moins de pièces)
en utilisant des pièces dont les valeurs possibles sont indiquées en paramètre de la fonction.

# Entrées
- somme : la somme d'argent à constituer avec les pièces
- valeursPieces : un tableau des valeurs de pièces existantes

# Sorties
- pieces : un ensemble de pièces dont la somme vaut somme
- err : une erreur qui doit valoir errImpossible si la somme ne peut pas être réalisée à partir des valeurs de pièces considérées.

# Exemples
meilleurDecomposition(15, []int{7, 8, 9}) = [8 7]
meilleurDecomposition(15, []int{2, 4, 6}) = errImpossible
*/

func decreasehistory(array []int, maxindex int) []int {
	fmt.Println("Problem ?")
	if array[len(array)-2] > 0 {
		array[len(array)-2] = array[len(array)-2] - 1
		array[len(array)-1] = maxindex
	} else {
		fmt.Println("inferior")
		for x := 3; x < len(array); x++ {
			if array[len(array)-x] > 0 {
				array[len(array)-x] = array[len(array)-x] - 1
				for y := 1; y < x; y++ {
					array[len(array)-y] = maxindex
				}
				return array
			}
		}

	}
	return array
}

func decreaseLastElement(history []int) []int {
	for i := 0; i < len(history); i++ {
		if i == len(history)-1 {
			history[i]--
		}
	}
	return history
}

func increaseLastElement(history []int) []int {
	for i := 0; i < len(history); i++ {
		if i == len(history)-1 {
			history[i]++
		}
	}
	return history
}

func getSomme(valeursPieces, history []int) int {
	var somme int = 0
	for i := 0; i < len(history); i++ {
		somme += valeursPieces[history[i]]
	}
	return somme
}

func sameSlice(array1, array2 []int) bool {
	if len(array1) != len(array2) {
		return false
	}
	for i := 0; i < len(array1); i++ {
		if array1[i] != array2[i] {
			return false
		}
	}
	return true
}

func removeLast(array []int) []int {
	var save []int
	for i := 0; i < len(array)-1; i++ {
		save = append(save, array[i])
	}
	return save
}

func convertFormat(history, valeursPieces []int) []int {
	var result []int
	for i := 0; i < len(history); i++ {
		result = append(result, valeursPieces[history[i]])
	}
	return result
}

func recur(history, valeursPieces, beforeHistory, beforeHistory2 []int, somme, index int) (pieces []int, err error) {
	var historySave []int = append(history, index)
	var historySave2 []int = beforeHistory
	fmt.Println(history)
	fmt.Println(getSomme(valeursPieces, history))
	if len(history) == 0 {
		if valeursPieces[index] < somme {
			history = append(history, index)
			return recur(history, valeursPieces, historySave, historySave2, somme, len(valeursPieces)-1)
		} else if valeursPieces[index] == somme {
			history = append(history, index)
			return history, nil
		} else {
			if index == 0 {
				return []int{}, errImpossible
			}
			return recur(history, valeursPieces, historySave, historySave2, somme, index-1)
		}
	} else {
		history = append(history, index)
		if sameSlice(beforeHistory, historySave) || sameSlice(historySave, beforeHistory2) {
			return []int{}, errImpossible
		}
		if getSomme(valeursPieces, history) < somme {
			if index == len(valeursPieces)-1 {
				return recur(history, valeursPieces, historySave, historySave2, somme, len(valeursPieces)-1)
			} else {
				history = removeLast(history)
				if getSomme(valeursPieces, append(history, index+1)) < somme {
					return recur(history, valeursPieces, historySave, historySave2, somme, index+1)
				} else {
					history = decreasehistory(history, len(valeursPieces)-1)
					return recur(history, valeursPieces, historySave, historySave2, somme, len(valeursPieces)-1)
				}
			}
		} else if getSomme(valeursPieces, history) > somme {
			if index == 0 {
				if len(history) > 2 {
					history = removeLast(history)
					history = decreasehistory(history, len(valeursPieces)-1)
					return recur(history, valeursPieces, historySave, historySave2, somme, len(valeursPieces)-1)
				} else {
					history[0] = history[0] - 1
					history = removeLast(history)
					return recur(history, valeursPieces, historySave, historySave2, somme, len(valeursPieces)-1)
				}
			} else {
				history = removeLast(history)
				return recur(history, valeursPieces, historySave, historySave2, somme, index-1)
			}
		} else {
			fmt.Println(convertFormat(history, valeursPieces))
			return convertFormat(history, valeursPieces), nil
		}
	}
}

func meilleurDecomposition(somme int, valeursPieces []int) (pieces []int, err error) {
	if somme == 0 {
		return []int{}, nil
	}
	if len(valeursPieces) > 0 {
		return recur([]int{}, valeursPieces, []int{}, []int{}, somme, len(valeursPieces)-1)
	}
	return nil, errImpossible
}
