package longueur

import "fmt"

/*
On souhaite trier un tableau de chaînes de caractères (minuscules sans accents) de la plus longue à la plus courte
(et par ordre alphabétique dans le cas où plusieurs chaînes ont la même longueur). La fonction trier doit réaliser ce tri en place.
On rappelle que l'opérateur < sur les chaînes de caractères définit l'ordre alphabétique.

# Entrée
- tab : le tableau de chaînes de caractères à trier

# Info
2022-2023, test3, exercice 1
*/

func decalLeft(tab *[]string, index int) {
	var result []string
	for i := 0; i < len(*tab); i++ {
		if i == index {
			result = append(result, "")
		}
		result = append(result, (*tab)[i])
	}
	*tab = result
}
func trier(tab *[]string) {
	var result []string = []string{}
	for i := 0; i < len(*tab); i++ {
		fmt.Println(result)
		var state bool = true
		for x := 0; x < len(result); x++ {
			if state {
				if len((*tab)[i]) > len(result[x]) {
					decalLeft(&result, x)
					result[x] = (*tab)[i]
					state = false
				} else if len((*tab)[i]) == len(result[x]) {
					if (*tab)[i] < result[x] {
						decalLeft(&result, x)
						result[x] = (*tab)[i]
						state = false
					} else if len(result) > x+1 {
						if !((*tab)[i] < result[x+1]) || !(len((*tab)[i]) == len(result[x+1])) {
							decalLeft(&result, x+1)
							result[x+1] = (*tab)[i]
							state = false
						}
					}
				}
			}
		}
		if state && len(result) != 0 {
			result = append(result, (*tab)[i])
		}
		if len(result) == 0 {
			result = append(result, (*tab)[i])
		}
	}
	fmt.Println(result)
	*tab = result
}
