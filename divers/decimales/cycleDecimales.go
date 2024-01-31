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

func trouveCycle(x, y uint) (cycle uint, err error) {
	flotant := fmt.Sprintf("%v", float64(x)/float64(y))
	fmt.Println(flotant, x, y)
	if len(strings.Split(flotant, ".")) == 2 {
		flotant = strings.Split(flotant, ".")[1]
	} else {
		flotant = "0"
	}
	var deleted int = 1
	if flotant[0] == '0' {
		deleted--
	}
	flotant = flotant[:len(flotant)-1]
	init := flotant
	var repeating bool = true
	for len(flotant) > 0 {
		var result []byte = []byte{flotant[0]}
		fmt.Println(flotant)
		for i := 1; i < len(flotant); i++ {
			fmt.Println(result, i)
			if flotant[i] == result[0] {
				var state bool = true
				for x := 0; x < len(result) && x+i < len(flotant); x++ {
					fmt.Println(string(result[x]), string(flotant[i+x]))
					if result[x] != flotant[i+x] {
						state = false
					}
				}
				if !state {
					result = append(result, flotant[i])
				} else {
					i = i + len(result) - 1
				}
			} else {
				fmt.Println("ryoki tenkai")
				result = []byte(flotant[:i+1])
			}
		}
		if len(result) != len(flotant) || !repeating {
			if len(flotant) >= 18 {
				return 0, errCycleTropLong
			}
			value, _ := strconv.Atoi(string(result))
			return uint(value), nil
		} else if deleted >= 3 {
			repeating = false
			flotant = init
		}
		flotant = flotant[1:]
		deleted++
	}
	return 0, errCycleTropLong
}
