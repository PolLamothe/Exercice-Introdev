package chiffres

import (
	"fmt"
	"strconv"
	"strings"
)

/*
1^1 + 2^2 + 3^3 + ... + 10^10 = 10405071317
Les 10 derniers chiffres de ce nombre sont 0405071317, qu'on représentera par
l'entier 405071317
La fonction chiffres doit donner l'entier qui représente, sur le même modèle,
les 10 derniers chiffres de 1^1 + 2^2 + ... + n^n pour un entier n donné en
argument.

Une fonction qui calcule chiffres(10000) en moins d'une minute sur ma machine
rapportera plus de points

# Entrée
- n : l'entier maximum utilisé dans le calcul

# Sorties
- c : l'entier représentant les 10 derniers chiffres de 1^1 + 2^2 + ... + n^n

Inspiré du problème 48 de projecteuler.net

# Info
2021-2022, test2, exercice 9
*/

func chiffres(n int64) (c int) {
	var somme int64 = 0
	var i int64
	for i = 1; i <= n; i = i + 1 {
		var temp int64 = i
		var x int64
		for x = 1; x < i; x = x + 1 {
			temp = (temp * i) % 10000000000
		}
		somme = (somme + temp) % 10000000000
	}
	var sommeString string = strconv.Itoa(int(somme))
	sommeString2 := strings.Split(sommeString, "")
	var temp []string
	for i := 10; i > 0; i = i - 1 {
		if len(sommeString2) >= i {
			temp = append(temp, sommeString2[len(sommeString2)-i])
		}
	}
	result, err := strconv.Atoi(strings.Join(temp, ""))
	fmt.Println(result)
	if err == nil {
		return result
	}
	return 0
}
