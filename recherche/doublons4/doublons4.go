package doublons4

import "fmt"

/*
On dispose d'un tableau d'entiers de longueur n et on suppose qu'il contient
exactement k fois chaque nombre compris entre 1 et n/k. On voudrait vérifier
cela, mais on ne connaît pas k. La fonction doublons va calculer k, si c'est
possible.

# Entrée
- tab : un tableau d'entiers

# Sortie
- k : un entier (s'il existe) tel que tab contient exactement k fois chaque
entiers de 1 à len(tab)/k
- ok : un booléen qui indique si k existe ou non
*/

func doublons(tab []int) (k int, ok bool) {
	if len(tab) == 0 {
		return 0, true
	}
	var count [][]int = [][]int{}
	for i := 0; i < len(tab); i++ {
		var state bool = false
		for x := 0; x < len(count); x++ {
			if count[x][0] == tab[i] {
				count[x][1]++
				state = true
			}
		}
		if !state {
			count = append(count, []int{tab[i], 1})
		}
	}
	for i := 1; i < len(count); i++ {
		if count[i][1] != count[i-1][1] {
			return 0, false
		}
	}
	var k_ int = count[0][1]
	fmt.Println(k_, count)
	if len(count) != len(tab)/k_ {
		return 0, false
	}
	for i := 1; i <= len(tab)/k_; i++ {
		var state bool = false
		for x := 0; x < len(count); x++ {
			if count[x][0] == i {
				state = true
			}
		}
		if state == false {
			return 0, false
		}
	}
	return k_, true
}
