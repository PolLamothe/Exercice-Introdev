package deuxpetits

/*
La fonction lesDeuxPlusPetits retourne les deux plus petits entiers présents dans un tableau

# Entrée
- t : un tableau d'entier qui en contient au moins 2

# Sorties
- v1 : le plus petit entier dans t
- v2 : le deuxième plus petit entier dans t

# Info
2022-2023, test 4, exercice 3
*/

func lesDeuxPlusPetits(t []int) (v1, v2 int) {
	var v1_ int
	var v2_ int
	var temp int
	for i := 0; i < len(t); i = i + 1 {
		if v1_ == 0 {
			v1_ = t[i]
		} else if v2_ == 0 {
			v2_ = t[i]
		}
		if t[i] <= v2_ {
			v2_ = t[i]
			if v2_ < v1_ {
				temp = v1_
				v1_ = v2_
				v2_ = temp
			}
		}
	}
	return v1_, v2_
}
