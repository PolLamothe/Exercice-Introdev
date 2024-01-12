package huitreines

import "fmt"

/*
Le problème des huit reines consiste à placer, sur un échiquier (un tableau de 8 cases par 8 cases), 8 reines, de telle sorte
qu'aucune d'entre-elles ne soit en position d'en manger une autre (c'est à dire de telle sorte qu'il n'y ait pas deux reines sur la
même ligne, la même colonne ou la même diagonale de l'échiquier).

Ce problème se généralise en problème des n reines. On cherche alors à placer de la même façon n reines sur un échiquier de n cases
par n cases.

Quand on analyse le problème, on se rend rapidement compte que chaque reine doit obligatoirement être sur une ligne différente des
autres : il y a exactement une reine par ligne de l'échiquier. Partant de là, on peut procéder récursivement pour résoudre le
problème :
1. placer une reine sur la première case de la première ligne (l) sans reine
2. si cette reine n'est pas mangée par une autre résoudre le même problème pour n-1 reines sur un échiquier ou les lignes 0 à l
sont occupées
3. s'il y a une solution, c'est gagné
4. sinon (si la reine est mangée ou si le problème à n-1 reines n'a pas de solution), placer la reine sur la case suivante de la
ligne (l) et recommencer à l'étape 2.

L'objectif de cet exercice est d'écrire une fonction huitreines qui résout le problème des n reines pour n donné.

# Entrée
- n : le nombre de reines à placer sur un échiquier de n cases par n cases

# Sortie
- plateau : un tableau de tableaux d'entiers de taille n*n qui contient des 0 là où il n'y a pas de reine et des 1 là où il y a
des reines dans une solution du problème des n reines (si elle existe)
- ok : un booléen qui indique si le problème a une solution ou pas

# Exemple
huitreines(4) = [[0 1 0 0] [0 0 0 1] [1 0 0 0] [0 0 1 0]] (solution non unique) (c'est horizontale)
*/

func getRowAndColumn(plateau [][]int, index int) (row, column int) {
	var columnCount int = 0
	var rowCount int = index
	for i := 0; i < len(plateau[index]); i++ {
		if plateau[index][i] == 1 {
			columnCount = i
		}
	}
	return rowCount, columnCount
}

func checkCoordonate(plateau [][]int, row, column int) bool {
	if plateau[row][column] == 1 {
		return true
	}
	return false
}

func checkColumn(plateau [][]int, index int) bool { //retourne false si il y'a une reine sur la colone
	row, column := getRowAndColumn(plateau, index)
	for i := 0; i < len(plateau); i++ {
		if plateau[i][column] == 1 && i != row {
			return false
		}
	}
	return true
}

func checkDiag(plateau [][]int, index int) bool { // retourn false si il y'a quelqu'un dans la diago
	row, column := getRowAndColumn(plateau, index)
	var yessir bool = true
	i := 1
	for yessir {
		if row-i >= 0 && column-i >= 0 {
			if checkCoordonate(plateau, row-1, column-1) {
				return false
			}
		}
		if row+i < len(plateau) && column-i >= 0 {
			if checkCoordonate(plateau, row+1, column-1) {
				return false
			}
		}
		if row-i >= 0 && column+i < len(plateau[0]) {
			if checkCoordonate(plateau, row-1, column+1) {
				return false
			}
		}
		if row+i < len(plateau) && column+i < len(plateau[0]) {
			if checkCoordonate(plateau, row+1, column+1) {
				return false
			}
		}
		i++
		if i == len(plateau) {
			yessir = false
		}
	}
	return true
}

func reverse(plateau [][]int) {

}

func recur(plateau [][]int, n, index, count int) (result [][]int, ok bool) {
	fmt.Println(plateau)
	var current []int = make([]int, n)
	if n == 1 {
		return [][]int{{1}}, true
	} else {
		if count == 0 {
			current[1] = 1
			plateau = append(plateau, current)
			return recur(plateau, n, 1, count+1)
		} else {
			if count == n {
				return plateau, true
			}
			var valid bool = false
			//for i := 0;i<
			for i := 0; i < len(plateau[0]); i++ {
				var tempPlateau [][]int = plateau
				var add []int = make([]int, n)
				add[i] = 1
				tempPlateau = append(tempPlateau, add)
				if checkColumn(tempPlateau, index) && checkDiag(tempPlateau, index) {
					plateau = tempPlateau
					valid = true
					return recur(plateau, n, index+1, count+1)
				}
			}
			if !valid {
				return nil, false
			}
		}
	}
	return nil, false
}

func huitreines(n int) (plateau [][]int, ok bool) {

	return recur([][]int{}, n, 0, 0)
}
