package suiteparam

/*
On considère êun ensemble de suites définies par U(n) = 2 * U(n-1) + U(0), chaque U(0) possible définissant une suite de cet ensemble. La fonction termeparam doit calculer le n-ième terme de la suite telle que U(0) = n.

Pour cet exercice, les boucles for sont interdites.

# Entrée
- n, un entier

# Sortie
- un, le terme U(n) de la suite telle que U(0) = n

# Info
2022-2023, test 2, exercice 6
*/

func recur(n0, n int) int {
	if n == 0 {
		return n0
	} else {
		return 2*recur(n0, n-1) + n0
	}
}

func termeparam(n int) (un int) {
	return recur(n, n)
}
