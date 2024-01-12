package deuxGrands

/*
La fonction lesDeuxPlusGrands retourne les deux plus grands entiers présents dans un tableau

# Entrée
- t : un tableau d'entiers qui en contient au moins 2

# Sorties
- v1 : le plus grand entier dans t
- v2 : le deuxième plus grand entier dans t

# Info
2022-2023, test 1, exercice 8
*/

func lesDeuxPlusGrands(t []int) (v1, v2 int) {
	var first int = 0
	var second int = 0
	for i := 0; i < len(t); i++ {
		if t[i] > first {
			var temp int = first
			first = t[i]
			second = temp
		} else if t[i] > second {
			second = t[i]
		}
	}
	return first, second
}
