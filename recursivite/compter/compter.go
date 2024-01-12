package compter

import (
	"strconv"
	"strings"
)

/*
La fonction compter(n, m) doit retourner une chaîne de caractère qui liste tous
les entiers dans l'ordre croissant de n à m puis dans l'ordre décroissant de m
à n, séparés par des signes "+". (Si n > m, ceci n'est pas possible et la fonction
retournera donc une chaîne vide.)

Vous ne devez pas utiliser de boucles, la fonction compter sera donc récursive.

# Entrées
- n : un entier
- m : un entier

# Sortie
- s : une chaîne de caractères

# Exemples
compter(5, 5) = "5"
compter(2, 5) = "2+3+4+5+4+3+2"

# Info
2021-2022, test3, exercice 4
*/

func strToArray(text string) []string {
	var result []string
	for i := 0; i < len(text); i++ {
		result = append(result, string(text[i]))
	}
	return result
}

func JoinArray(text1, text2 []string) []string {
	for i := 0; i < len(text2); i++ {
		text1 = append(text1, text2[i])
	}
	return text1
}

func reverse(array []string) []string {
	var result []string
	for i := len(array) - 1; i > -1; i-- {
		result = append(result, array[i])
	}
	return result
}

func compter(n, m uint) (s string) {
	if n > m {
		return ""
	}
	if n == m {
		return strings.Join(strToArray(strconv.Itoa(int(n))), "")
	}
	if n+1 == m {
		return strings.Join([]string{strconv.Itoa(int(n)), "+", strconv.Itoa(int(m)), "+", strconv.Itoa(int(n))}, "")
	}
	var brefore []string = strToArray(compter(n, m-1))
	var brefore1 []string = reverse(brefore)
	return strings.Join(JoinArray(append(JoinArray(brefore[:int(float64(m-n)*1.75)], []string{"+"}), strconv.Itoa(int(m))), JoinArray([]string{"+"}, brefore1[int(float64(m-n)*1.75)-1:])), "")
}
