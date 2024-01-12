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

func recur(n, nInit int) (un int) {
	if n == 0 {
		return nInit
	}
	return 3*recur(n-1, nInit) - n*recur(0, nInit)
}

func termeparam(n int) (un int) {
	return recur(n, n)
}
