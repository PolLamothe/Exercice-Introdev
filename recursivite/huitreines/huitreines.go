package huitreines

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

func checkDiag(plat [][]int, n, x, y int) bool {
	for i := -n; i < n; i++ {
		if i != 0 {
			if x+i >= 0 && x+i < n {
				if y+i >= 0 && y+i < n {
					if plat[y+i][x+i] == 1 {
						return false
					}
				}
				if y-i >= 0 && y-i < n {
					if plat[y-i][x+i] == 1 {
						return false
					}
				}
			}
			if x-i >= 0 && x-i < n {
				if y+i >= 0 && y+i < n {
					if plat[y+i][x-i] == 1 {
						return false
					}
				}
				if y-i >= 0 && y-i < n {
					if plat[y-i][x-i] == 1 {
						return false
					}
				}
			}
		}
	}
	return true
}

func isRight(plateua [][]int, n int) bool {
	var plat = append([][]int{}, plateua...)
	for i := 0; i < len(plat); i++ {
		var queenCount = 0
		if len(plat[i]) != n {
			return false
		}
		for x := 0; x < len(plat[i]); x++ {
			if plat[i][x] == 1 {
				queenCount++
			}
		}
		if queenCount != 1 {
			return false
		}
	}
	for i := 0; i < len(plat); i++ {
		var queenCount int = 0
		for x := 0; x < len(plat[i]); x++ {
			if plat[x][i] == 1 {
				queenCount++
			}
		}
		if queenCount != 1 {
			return false
		}
	}
	for i := 0; i < n && len(plat) > 1; i++ {
		for x := 0; x < n; x++ {
			if plat[i][x] == 1 {
				if !checkDiag(plat, n, x, i) {
					return false
				}
			}
		}
	}
	return true
}

func recur(previous [][]int, n int) (plateau [][]int, ok bool) { //je ne peux pas utiliser de "vrai" fonction récurcive car sinon il y'a un overflow a partir de n = 9
	var plat = append([][]int{}, previous...)
	for {
		for i := 0; i < n; i++ {
			var queenPos int = 0
			for x := 0; x < n; x++ {
				if plat[i][x] == 1 {
					queenPos = x
					break
				}
			}
			if queenPos < n-1 {
				var copy []int = append([]int{}, plat[i]...)
				copy[queenPos] = 0
				copy[queenPos+1] = 1
				plat[i] = copy
				break
			} else {
				var state bool = true
				for i := 0; i < n; i++ {
					if plat[i][n-1] != 1 {
						state = false
					}
				}
				if state {
					return plat, false
				}
				var copy []int = append([]int{}, plat[i]...)
				copy[queenPos] = 0
				copy[0] = 1
				plat[i] = copy
			}
		}
		if isRight(plat, n) {
			return plat, true
		}
	}
	return
}

func huitreines(n int) (plateau [][]int, ok bool) {
	if n == 1 {
		return [][]int{{1}}, true
	}
	var tab []int = []int{1}
	for i := 1; i < n; i++ {
		tab = append(tab, 0)
	}
	var tabtab [][]int
	for i := 0; i < n; i++ {
		tabtab = append(tabtab, tab)
	}
	return recur(tabtab, n)
}
