package occurrencesmax

/*
Étant donné un tableau d'entiers t, la fonction occurrencesmax doit retourner
le plus grand entier qui apparaît dans t et le nombre de fois que cet entier
apparaît. En cas d'égalité on choisira arbitrairement l'un des entiers
à égalité. Si le tableau est vide on retournera un entier quelconque et 0 pour son nombre d'apparitions.

# Entrées
- t : le tableau dans lequel chercher

# Sortie
- n : le plus grand entier qui apparaît dans t
- occ : le nombre de fois que n apparaît dans t

# Exemple
occurrencesmax([]int{1, 2, 3, 4, 3}) = 4, 1
*/
func occurrencesmax(t []int) (n int, occ int) {
	var count int = 0
	var max int
	for i := 0; i < len(t); i++ {
		if count == 0 {
			max = t[i]
			count++
		} else {
			if t[i] > max {
				count = 1
				max = t[i]
			} else if t[i] == max {
				count++
			}
		}
	}
	return max, count
}
