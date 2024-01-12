package cycleDecimales

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var errCycleTropLong error = errors.New("Le cycle de décimales en contient 18 ou plus")

/*
Lorsqu'on divise un nombre entier par un autre nombre entier, le résultat est toujours un nombre décimal dont la fin (une ou plusieurs décimales) se répète à l'infini.

Par exemple, 1/3 vaut 0,3333… avec un 3 qui se répète à l'infini, 1/2 vaut 0,5000… avec un 0 qui se répète à l'infini et 7/11 vaut 0,6363… avec 63 (deux décimales) qui se réoète à l'infini.

La fonction trouveCycle doit retourner un entier qui représente les décimales qui se répètent dans le résultat de la division d'un entier x par un entier y. Elle retournera toujours le moins de décimales possibles (par exemple pour 7/11 on retournera 63 et pas 6363) et les décimales représentant le cycle dès qu'il commence (par exemple pour 7/11 on retournera 63 et pas 36).

Il est possible que le décimales qui se répètent ne puissent pas se représenter avec un entier (car il y en a trop). On considérera que c'est le cas quand il y en a 18 ou plus. Dans ce cas, il faudra retourner l'erreur errCycleTropLong

# Entrées
- x : un entier positif ou nul
- y : un entier strictement positif

# Sortie
- cycle : un entier représentant les décimales qui se répètent dans x/y s'il y en a strictement moins de 18
- err : une erreur qui vaut nil s'il y a moins de 18 décimales qui se répètent et  qui vaut errCycleTropLong s'il y en a 18 ou plus

# Exemple
trouveCycle(1, 3) = 3
trouveCycle(1, 2) = 0
trouveCycle(7, 11) = 63

# Info
2022-2023, test 1, exercice 9 (dans le test il n'était pas demandé de traiter les cas d'erreur)
*/

func recur(init, treated []string, index, count int) bool {
	if init[index] == treated[count] {
		if len(treated) == count+1 {
			return true
		} else {
			return recur(init, treated, index+1, count+1)
		}
	} else {
		return false
	}

}

func deleteLast(array []string) []string {
	var result []string
	for i := 0; i < len(array)-1; i++ {
		result = append(result, array[i])
	}
	return result
}

func trouveCycle(x, y uint) (cycle uint, err error) {
	var number float64 = float64(x) / float64(y)
	var string2 string = strconv.FormatFloat(number, 'f', -1, 64)
	var string3 []string = strings.Split(string2, "")
	var treated []string
	fmt.Println(string3)
	for x := 0; x < len(string3); x++ {
		for i := 2; i < len(string3)-x; i++ {
			fmt.Println(treated)
			i = i + x
			if len(treated) == 0 {
				if string3[i] != "0" {
					treated = append(treated, string3[i])
				}
			} else {
				if string3[i] == treated[0] {
					fmt.Println(string3[i])
					fmt.Println(recur(string3, treated, i, 0))
					if recur(string3, treated, i, 0) {
						for j := len(treated) - 1; j >= 0; j-- {
							if treated[j] == "0" {
								treated = deleteLast(treated)
							} else {
								break
							}
						}
						value, _ := strconv.ParseUint(strings.Join(treated, ""), 10, 64)
						return uint(value), nil
					} else {
						treated = append(treated, string3[i])
					}
				} else {
					treated = append(treated, string3[i])
				}
			}
		}
		if x == 18 {
			return 0, errCycleTropLong
		}
		treated = []string{}
	}
	return 0, errCycleTropLong
}
