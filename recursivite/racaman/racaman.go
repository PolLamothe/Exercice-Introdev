package racaman

import "fmt"

/*
La suite de Racaman est définie par a(1) = 1, puis pour n > 1 par :
- a(n-1) - n si ce nombre est strictement supérieur à 0 et n'a encore jamais été
vu dans la suite
- a(n-1) + n sinon

La fonction racaman doit calculer a(n) pour n supérieur ou égal à 1.

# Entrée
- n : le numéro du terme de la suite à calculer

# Sortie
- an : la valeur du terme de la suite calculé (si ce terme n'est pas défini, on
retournera -1)

# Exemple
racaman(4) = 2
*/

func haveBeenSeen(array []int, number int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == number {
			return true
		}
	}
	return false
}

func recur(current, value, n int, history []int) int {
	history = append(history, value)
	fmt.Println(value)
	if n == current {
		return value
	}
	if value-(current+1) > 0 && !haveBeenSeen(history, value-(current+1)) {
		return recur(current+1, value-(current+1), n, history)
	} else {
		return recur(current+1, value+current+1, n, history)
	}
}

func racaman(n int) (an int) {
	if n <= 0 {
		return -1
	}
	var history []int
	return recur(1, 1, n, history)
}
