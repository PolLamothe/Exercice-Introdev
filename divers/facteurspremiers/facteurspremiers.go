package facteurspremiers

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

func mypremiers(n int) (p []int) {
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
	return result
}

func facteursPremiers(n uint) (facteurs []uint) {
	var premier []int = mypremiers(int(n) + 1)
	var reste uint = n
	var result []uint = []uint{}
	var state bool = false
	for reste > 1 {
		for i := len(premier) - 1; i >= 0 && state; i-- {
			if reste%uint(premier[i]) == 0 {
				result = append(result, uint(premier[i]))
				reste /= uint(premier[i])
				state = false
			}
		}
		state = true
	}
	return result
}
