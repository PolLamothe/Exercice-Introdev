package nombrespremiers

import "fmt"

/*
La fonction premiers doit retourner un tableau contenant les nombres
premiers compris entre 0 et un entier n.

# Entrée
- n : un nombre entier

# Sortie
- p : un tableau contenant tous les nombres premiers compris entre 0 et n inclus, s'il n'existe pas de tels nombres premiers, il faut que p soit un tableau contenant 0 éléments

# Exemple
premiers(10) = [2 3 5 7] (l'ordre n'a pas d'importance)
*/
func premiers(n int) (p []int) {
	var result []int = []int{}
	for i := 2; i < n; i++ {
		var state bool = true
		for x := 2; x < i; x++ {
			if state {
				if i%x == 0 {
					state = false
				}
			}
		}
		if state {
			result = append(result, i)
		}
	}
	fmt.Println(result)
	return result
}
