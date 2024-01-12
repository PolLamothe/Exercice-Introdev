package suitecond

/*
On considère la suite définie par
- U(n) = 2*U(n-1) + 1 si U(n-1) < 100
- U(n) = U(n-1) - 100 si U(n-1) >= 100
- U(0) = 2.
La fonction terme doit permettre d'obtenir le n-ième terme de cette suite.

Pour cet exercice, les boucles for sont interdites

# Entrée
- n, un entier

# Sortie
- un, le terme U(n) de la suite

# Info
2022-2023, test 2, exercice 1
*/

func recur(n, index, value int) int {
	if n == index {
		return value
	}
	if value < 100 {
		return recur(n, index+1, 2*value+1)
	} else {
		return recur(n, index+1, value-100)
	}
}

func terme(n uint) (un int) {
	return recur(int(n), 0, 2)
}
