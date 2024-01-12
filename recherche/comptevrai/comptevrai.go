package comptevrai

/*
La fonction compteVrai doit indiquer combien de fois la valeur true apparaît dans un tableau de booléens.

# Entrée
- t : un tableau de booléens

# Sortie
- nombre : le nombre de fois que la valeur true apparaît dans t

# Info
2023-2024, test 1, exercice 2
*/

func compteVrai(t []bool) (nombre int) {
	if len(t) == 0 {
		return 0
	}
	if len(t) == 1 {
		if t[0] == true {
			return 1
		} else {
			return 0
		}
	}
	var counter int = 0
	if t[0] == true {
		counter = 1
	}
	return compteVrai(t[1:]) + counter
}
