package suiteparam

/*
On considère un ensemble de suites définies par U(n) = 3 * U(n-1) - n * U(0), chaque U(0) possible définissant une suite de cet ensemble. La fonction termeparam doit calculer le n-ième terme de la suite telle que U(0) = n.

Pour cet exercice, les boucles for sont interdites.

# Entrée
- n, un entier

# Sortie
- un, le terme U(n) de la suite telle que U(0) = n

# Info
2022-2023, test 4, exercice 9
*/

func recur2(array []int, n, index int) (nValue int) {
	if index <= n {
		array = append(array, 3*array[index-1]-index*array[0])
		return recur2(array, n, index+1)
	} else {
		return array[n]
	}
}

func recur(n int) int {
	var array []int = []int{n}
	var nValue int = recur2(array, n, 1)
	return nValue
}

func termeparam(n int) (un int) {
	return recur(n)
}
