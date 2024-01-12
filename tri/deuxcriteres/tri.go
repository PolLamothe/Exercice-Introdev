package tri

/*
La fonction tri doit trier un tableau d'entiers de la manière suivante : on trouvera d'abord tous les nombres strictement négatifs du
tableau, dans l'ordre décroissant, puis tous les nombres positifs ou nuls, dans l'ordre croissant.

# Entrée
- t, le tableau à trier (en place)

# Info
2022-2023, test 4, exercice 8
*/

func decalLeft(t *[]int, index int) {
	var copy []int = *t
	var length = len(copy)
	var temp int = (copy)[length-1]
	var temp2 int
	var temp3 int
	for i := 0; i < length; i = i + 1 {
		if i == index {
			var isOdd bool = index%2 == 0
			if isOdd {
				temp2 = (*t)[i]
			} else {
				temp3 = (*t)[i]
			}
			(*t)[i] = 0
		}
		if i > index {
			if i%2 == 0 {
				temp2 = (*t)[i]
				(*t)[i] = temp3
			} else {
				temp3 = (*t)[i]
				(*t)[i] = temp2
			}
		}
		if i == length-1 {
			(*t) = append((*t), temp)
			return
		}
	}
}

func tri(t []int) []int {
	var tabNeg []int
	var tabPos []int
	for i := 0; i < len(t); i = i + 1 {
		if t[i] >= 0 {
			if len(tabPos) == 0 {
				tabPos = append(tabPos, t[i])
			} else {
				var clone []int = tabPos
				for x := 0; x < len(tabPos); x = x + 1 {
					if clone[x] > t[i] {
						decalLeft(&clone, x)
						clone[x] = t[i]
					} else if clone[x] < t[i] && x == len(clone)-1 {
						clone = append(clone, t[i])
					}
				}
				tabPos = clone
			}
		} else {
			if len(tabNeg) == 0 {
				tabNeg = append(tabNeg, t[i])
			} else {
				var clone []int = tabNeg
				for x := 0; x < len(tabNeg); x = x + 1 {
					if clone[x] < t[i] {
						decalLeft(&clone, x)
						clone[x] = t[i]
					} else if clone[x] > t[i] && x == len(clone)-1 {
						clone = append(clone, t[i])
					}
				}
				tabNeg = clone
			}
		}
	}
	t = []int{}
	for i := 0; i < len(tabNeg); i = i + 1 {
		t = append(t, tabNeg[i])
	}
	for i := 0; i < len(tabPos); i = i + 1 {
		t = append(t, tabPos[i])
	}
	return t
}
