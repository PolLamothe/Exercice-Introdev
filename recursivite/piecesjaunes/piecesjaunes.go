package piecesjaunes

import (
	"errors"
	"fmt"
)

var errImpossible error = errors.New("impossible de constituer cette somme avec ces pièces")

/*
Les pièces et billets que nous utilisons dans la vie de tous les jours ont des valeurs qui ont été choisies pour qu'il soit simple de rendre la monnaie (il suffit de toujours donner la pièce ou le billet de plus grande valeur possible, jusqu'à avoir rendu la somme voulue). Dans cet exercice, on imagine que ce n'est pas le cas : on dispose d'un ensemble de pièces de valeurs quelconques et on cherche à constituer une somme avec le moins de pièces possible.

La fonction meilleurDecomposition doit trouver la meilleur décomposition d'une somme d'argent (celle qui utilise le moins de pièces) en piochant des pièces dans l'ensemble donné en argument de la fonction.

# Entrées
- somme : la somme d'argent à constituer avec les pièces
- piecesDisponibles : un tableau des pièces disponibles

# Sorties
- pieces : un ensemble de pièces dont la somme vaut somme
- err : une erreur qui doit valoir errImpossible si la somme ne peut pas être réalisée à partir des pièces considérées.

# Exemples
meilleurDecomposition(15, []int{2, 3, 5, 5, 5}) = [5 5 5]
meilleurDecomposition(15, []int{2, 3, 5}) = errImpossible
*/

func getSupply(supply [][]int, value int) int { //obtenir le nombre de pieces disponible
	for i := 0; i < len(supply); i++ {
		if supply[i][0] == value {
			return supply[i][1]
		}
	}
	return 0
}

func getNumber(tab []int, value int) int { //obtenir le nombre d'élément précis dans une array
	var result int = 0
	for i := 0; i < len(tab); i++ {
		if tab[i] == value {
			result++
		}
	}
	return result
}

func getSum(tab []int) int { //obtenir la somme des éléments d'une array
	var result int = 0
	for i := 0; i < len(tab); i++ {
		result += tab[i]
	}
	return result
}

func getIndex(tab []int, value int) int { //obtenir l'index d'un élément dans une array en partant du début
	for i := 0; i < len(tab); i++ {
		if tab[i] == value {
			return i
		}
	}
	return -1
}

func getIndexFromSupply(tab [][]int, value int) int {
	for i := 0; i < len(tab); i++ {
		if tab[i][0] == value {
			return i
		}
	}
	return -1
}

func decreaseCurrent(current []int, supply [][]int) []int {
	var index int = getIndexFromSupply(supply, current[len(current)-1])
	for i := index - 1; i > -1; i-- {
		if getNumber(current, supply[i][0]) < getSupply(supply, supply[i][0]) {
			current[len(current)-1] = supply[i][0]
			return current
		}
	}
	return []int{}
}

func deleteFirst(tab []int) []int {
	var result []int = []int{}
	for i := 1; i < len(tab); i++ {
		result = append(result, tab[i])
	}
	return result
}

func recur(supply [][]int, current []int, somme, value int) (pieces []int, err error) {
	fmt.Println(current, value)
	if getSupply(supply, value) <= getNumber(current, value) {
		if value == supply[0][0] {
			return []int{}, errImpossible
		}
		return recur(supply, current, somme, supply[getIndexFromSupply(supply, value)-1][0])
	}
	if value+getSum(current) < somme {
		if getNumber(current, value) <= getSupply(supply, value) {
			current = append(current, value)
			return recur(supply, current, somme, supply[len(supply)-1][0])
		} else {
			for i := len(supply) - 1; i < -1; i-- {
				if getNumber(current, supply[i][0]) < supply[i][1] {
					if getSum(current) <= somme {
						return recur(supply, current, somme, supply[i][0])
					}
				}
			}
			return []int{}, errImpossible
		}
	} else if value+getSum(current) > somme {
		if getIndexFromSupply(supply, value) > 0 {
			return recur(supply, current, somme, supply[getIndexFromSupply(supply, value)-1][0])
		} else if len(current) > 1 {
			current = decreaseCurrent(current, supply)
			return recur(supply, current, somme, supply[len(supply)-1][0])
		} else if current[0] != supply[0][0] {
			if len(current) > 1 {
				return recur(supply, deleteFirst(current), somme, supply[len(supply)-1][0])
			} else if current[0] != supply[0][0] {
				current = decreaseCurrent(current, supply)
				return recur(supply, current, somme, supply[len(supply)-1][0])
			} else {
				return []int{}, errImpossible
			}
		} else {
			return []int{}, errImpossible
		}
	} else {
		current = append(current, value)
		return current, nil
	}
}

func meilleurDecomposition(somme int, piecesDisponibles []int) (pieces []int, err error) {
	var supply [][]int = [][]int{}
	if somme == 0 {
		return []int{}, nil
	}
	if len(piecesDisponibles) == 0 {
		return []int{}, errImpossible
	}
	for i := 0; i < len(piecesDisponibles); i++ {
		var state bool = false
		for x := 0; x < len(supply); x++ {
			if supply[x][0] == piecesDisponibles[i] {
				state = true
				supply[x][1]++
			}
		}
		if !state {
			supply = append(supply, []int{piecesDisponibles[i], 1})
		}
	}
	return recur(supply, []int{}, somme, piecesDisponibles[len(piecesDisponibles)-1])
}
