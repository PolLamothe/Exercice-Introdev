package vraifaux

import "fmt"

/*
On souhaite trier un tableau de booléens en mettant ceux dont la valeur est true au début et ceux dont la valeur est false à la fin. La fonction trier doit réaliser ce tri en place.

# Entrée
- tab : le tableau de booléens à trier

# Info
2022-2023, test3, exercice 0
*/

func decalLeft(array *[]bool, index int) {
	var result []bool = []bool{}
	for i := 0; i < len(*array); i++ {
		if i == index {
			result = append(result, true)
		}
		result = append(result, (*array)[i])
	}
	*array = result
}

func trier(tab *[]bool) {
	var result []bool
	for i := 0; i < len(*tab); i++ {
		fmt.Println(result)
		if i == 0 {
			if (*tab)[i] == true {
				result = append(result, true)
			} else {
				result = append(result, false)
			}
		} else {
			if (*tab)[i] == true {
				decalLeft(&result, 0)
			} else {
				result = append(result, false)
			}
		}
	}
	*tab = result
}
