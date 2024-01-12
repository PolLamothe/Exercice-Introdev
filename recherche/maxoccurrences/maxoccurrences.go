package maxoccurrences

/*
Étant donné un tableau d'entiers t, la fonction maxoccurrences doit retourner
l'entier qui apparaît le plus souvent dans t et le nombre de fois que cet
entier apparaît. En cas d'égalité on choisira arbitrairement l'un des entiers
à égalité. Si le tableau est vide on retournera un entier quelconque et 0 pour son nombre d'apparitions.

# Entrées
- t : le tableau dans lequel chercher

# Sortie
- n : l'entier qui apparaît le plus de fois dans t
- occ : le nombre de fois que n apparaît dans t

# Exemple
maxoccurrences([]int{1, 2, 3, 4, 3}) = 3, 2
*/
func maxoccurrences(t []int) (n int, occ int) {
	var count [][]int = [][]int{}
	for i := 0; i < len(t); i++ {
		var state bool = false
		for x := 0; x < len(count); x++ {
			if count[x][0] == t[i] {
				count[x][1]++
				state = true
			}
		}
		if !state {
			count = append(count, []int{t[i], 1})
		}
	}
	var number int = 0
	var counter int = 0
	for i := 0; i < len(count); i++ {
		if count[i][1] > counter {
			counter = count[i][1]
			number = count[i][0]
		}
	}
	return number, counter
}
