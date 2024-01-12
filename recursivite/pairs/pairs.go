package pairs

/*
La fonction pairs doit indiquer combien il y a de nombres pairs dans un tableau donné en paramètre. On peut remarquer que le nombre de nombres pairs dans un tableau t est la somme du nombre de nombres pairs dans la première moitié de t et du nombre de nombres pairs dans la deuxième moitié de t.

Vous ne devez pas utiliser de boucles for dans cet exercice.

# Entrée
- t, le tableau dans lequel compter les nombres pairs

# Sortie
- num, le nombre de nombres pairs dans t

# Info
2022-2023, test 2, exercice 3
*/

func recur(array []int, counter, index int) int {
	if index == len(array) {
		return counter
	}
	if array[index]%2 == 0 {
		counter++
	}
	return recur(array, counter, index+1)
}
func pairs(t []int) (num int) {
	return recur(t, 0, 0)
}
