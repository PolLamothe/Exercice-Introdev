package parentheses

import "fmt"

/*
Un bon parenthésage est une chaîne de caractères constituée uniquement de parenthèses ouvrantes et fermantes et telle que
- toute parenthèse ouvrante est bien fermée plus tard dans la chaîne,
- toute parenthèse fermante a été bien ouverte plus tôt dans la chaîne.

Par exemple :
(()) est un parenthésage correct
()() est un parenthésage correct
()(( n'est pas un parenthésage correct
)(() n'est pas un parenthèsage correct

La fonction parentheses doit retourner tous les bons parenthésages de longueur n (si cela est possible).

On peut, pour cela, construire les parenthésages parenthése par parenthése (on passera donc par des parenthésages qui ne sont pas bons, mais dont on est certains qu'on pourra les rendre bons plus tard), en faisant attention de ne pas ouvrir une parenthèse quand ce n'est pas possible (si on ne pourra jamais la fermer) et de ne pas fermer de parenthèses qui n'ont pas été ouvertes.

# Entrée
- n, la longueur des parenthésages souhaités

# Sorties
- sequences, l'ensemble des bon parenthésages de longueur n s'il en existe (une valeur quelconque sinon), sans doublons, l'ordre des parenthésages dans le tableau n'a pas d'importance
- ok, un booléen qui vaut true s'il existe un ou des bons parenthésages de longueur n et qui vaut false sinon

# Info
2022-2023, test 2, exercice 9
*/

func IsCorrect(text string) bool {
	var openCount int = 0
	var closeCount int = 0
	for i := 0; i < len(text); i++ {
		if string(text[i]) == "(" {
			openCount++
		} else {
			closeCount++
		}
	}
	if openCount != closeCount {
		return false
	}
	var currentOpen int = 0
	for i := 0; i < len(text); i++ {
		if string(text[i]) == "(" {
			currentOpen++
		} else {
			currentOpen--
		}
		if currentOpen < 0 {
			return false
		}
	}
	return true
}

func recur(text string, result []string) []string {
	fmt.Println(text)
	if IsCorrect(text) {
		result = append(result, text)
	}
	bytes := []byte(text) // Convert text to []byte
	for i := 0; i < len(text); i++ {
		if string(bytes[i]) == "(" {
			bytes[i] = ')' // Modify the character at index i
			for x := 0; x < i; x++ {
				bytes[x] = '('
			}
			return recur(string(bytes), result)
		}
	}
	return result
}

func parentheses(n int) (sequences []string, ok bool) {
	if n%2 != 0 {
		return []string{}, false
	}
	if n == 0 {
		return []string{""}, true
	}
	var startString string = ""
	for i := 0; i < n; i++ {
		startString = startString + "("
	}
	return recur(startString, []string{}), true
}
