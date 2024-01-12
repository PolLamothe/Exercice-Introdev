package suitemoinssimple

import "fmt"

/*
On considère la suite dont le terme u(n) est défini par
- u(0)=1
- si n > 0 et 100 est divisible par u(n-1) alors u(n) vaut u(n-1) - 1
- sinon u(n) est le reste de la division de u(n-1) + 11 par 100
La fonction terme calcule les termes de cette suite.

# Entrée
- n : un entier indiquant le numéro du terme à calculer

# Sortie
- un : le terme u(n) de la suite

# Remarque
Les boucles ne sont pas autorisées pour résoudre cet exercice
À toute fin utile, on signale que 100 n'est pas divisible par 0
Votre fonction devra pouvoir calculer U(1000) en moins de 2 secondes

# Info
2022-2023, test 4, exercice 6
*/

func recur(array []int, n, f int) []int {
	if n-1 == f {
		return array
	} else {
		if n > 1 && (100%array[n-1] == 0) {
			array = append(array, array[n-1]-1)
		} else {
			array = append(array, (array[n-1]+11)%100)
		}
		fmt.Println(array)
		return recur(array, n+1, f)
	}
}

func terme(n int) (un int) {
	if n == 0 {
		return 1
	}
	var array []int = []int{0}
	array = recur(array, 1, n)
	return array[n-1]
}
