package suite

/*
On considère la suite (un) définie par u(n) = 3 * u(n-1) + 5 et u(0) = 0
La fonction suite doit calculer le n-ième terme de (un) de manière récursive :
une fonction qui n'est pas récursive rapportera moins de points.

# Entrée
- n : un entier indiquant le terme de la suite à calculer

# Sortie
- un : le n-ième terme de la suite

# Info
2021-2022, test 1, exercice 3
*/

func recur(n, index, value int) int {
	if index == n {
		return value
	}
	return recur(n, index+1, value*3+5)
}

func suite(n uint) (un uint) {
	return uint(recur(int(n), 0, 0))
}
