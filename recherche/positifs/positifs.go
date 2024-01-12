package positifs

/*
La fonction positifs doit indiquer combien il y a de nombres strictement positifs dans un tableau donné en paramètre. On peut remarquer que le nombre de nombres strictement positifs dans un tableau t est la somme du nombre de nombres strictement positifs dans la première moitié de t et du nombre de nombres strictement positifs dans la deuxième moitié de t.

Vous ne devez pas utiliser de boucles for dans cet exercice.

# Entrée
- t, le tableau dans lequel compter les nombres strictement positifs

# Sortie
- num, le nombre de nombres strictement positifs dans t

# Info
2022-2023, test 4, exercice 5
*/

func recur(t []int, counter, index int) int {
	if index < len(t) {
		if t[index] > 0 {
			counter = counter + 1
		}
		return recur(t, counter, index+1)
	} else {
		return counter
	}
}

func positifs(t []int) (num int) {
	return recur(t, 0, 0)
}
