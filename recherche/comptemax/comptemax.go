package comptemax

/*
La fonction compteMax doit trouver le plus grand entier dans un tableau et indiquer combien de fois cet entier apparaît dans le tableau.

# Entrée
- t : un tableau d'entiers

# Sortie
- val : le plus grand entier dans t
- nombre : le nombre de fois que val apparaît dans t

# Info
- 2023-2024, test 1, exercice 7
*/

func compteMax(t []int) (val, nombre int) {
	var counter int = 0
	var max int = 0
	for i := 0; i < len(t); i++ {
		if t[i] > max {
			counter = 1
			max = t[i]
		} else if t[i] == max {
			counter++
		}
	}
	return max, counter
}
