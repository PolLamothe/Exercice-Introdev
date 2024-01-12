package doublons3

import (
	"fmt"
)

/*
On dispose d'un tableau d'entiers de longueur n et on suppose qu'il contient
exactement k fois chaque nombre compris entre 1 et n/k. On voudrait vérifier
cela. C'est le travail de la fonction doublons.

# Entrée
- tab : un tableau d'entiers
- k : un entier

# Sortie
- ok : un booléen qui doit valoir true si tab contient exactement k fois chaque
entiers de 1 à len(tab)/k et false sinon
*/

func isInDataSlice(dataSlice []int, e int) bool {
	for i := 0; i < len(dataSlice); i = i + 1 {
		if dataSlice[i] == e {
			return true
		}
	}
	return false
}

func doublons(tab []int, k int) (ok bool) {
	var countSlice []int
	var dataSlice []int
	if len(tab) > 0 {
		for i := 0; i < len(tab); i = i + 1 {
			for x := 0; x <= len(dataSlice); x = x + 1 {
				if x == 0 {
					if i == 0 {
						countSlice = append(countSlice, 0)
						dataSlice = append(dataSlice, tab[i])
					}
				} else {
					if len(dataSlice) >= x {
						if dataSlice[x-1] == tab[i] {
							fmt.Println(tab[i])
							countSlice[x-1] = countSlice[x-1] + 1
						}
					}
					if !isInDataSlice(dataSlice, tab[i]) {
						countSlice = append(countSlice, 0)
						dataSlice = append(dataSlice, tab[i])
					}
				}
			}
		}
		fmt.Println(dataSlice)
		fmt.Println(countSlice)
		var goodArray []int
		for i := 1; i <= len(tab)/k; i = i + 1 {
			goodArray = append(goodArray, i)
		}
		for i := 1; i < len(tab)/k; i = i + 1 {
			if dataSlice[i] > goodArray[len(goodArray)-1] || dataSlice[i] < goodArray[0] {
				return false
			}
			for x := 0; x < len(dataSlice); x = x + 1 {
				if dataSlice[x] == i {
					if countSlice[x] != k {
						return false
					}
				}
			}
		}
	}
	return true
}
