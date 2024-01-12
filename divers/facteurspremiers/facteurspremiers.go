package facteurspremiers

import "fmt"

/*
La fonction facteursPremiers doit retourner un tableau contenant la liste de tous
les facteurs premiers d'un entier n, doublons compris

# Entrée
- n : un nombre entier positif

# Sortie
- facteurs : un tableau contenant tous les facteurs premiers de n, si n vaut 0 il
faut retourner un tableau à zéro éléments.

# Exemple
premiers(24) = [2 2 2 3] (l'ordre n'a pas d'importance)
*/

func multiply(array []uint) int {
	var result int = 1
	for i := 0; i < len(array); i++ {
		result *= int(array[i])
	}
	return result
}

func facteursPremiers(n uint) (facteurs []uint) {
	var i uint = 2
	if n == 1 {
		return []uint{}
	}
	for ; i <= n; i++ {
		if n%i == 0 {
			var x uint = 2
			var state bool = true
			for ; x < i; x++ {
				if i%x == 0 {
					state = false
				}
			}
			if state {
				fmt.Println(n, i)
				return append(facteursPremiers(n/i), i)
			}
		}
	}
	return []uint{}
}
