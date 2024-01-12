package puissant

import "fmt"

/*
On dit qu'un mot est puissant s'il est la répétition d'un même motif 2 ou plus
de fois. Ainsi, "bonbon" est puissant (le même motif se répète 2 fois, on dit
aussi que c'est un carré) et "lalala" est puissant (le même motif se répète 3
fois, on dit aussi que c'est un cube). Par contre "bonjour" n'est pas puissant,
ni "bonbons" (le s final casse la répétition du motif "bon").

La fonction puissant doit déterminer si un mot, donné sous la forme d'une chaîne
de caractères ne comprenant que des lettres minuscules non accentuées, est
puissant ou non.

Des points seront donnés pour la réalisation d'une fonction aussi efficace ou
plus efficace que celle du prof.

# Entrée
- mot : le mot à analyser

# Sortie
- estPuissant : true si mot est puissant et false sinon

# Info
2021-2022, test 1, exercice 9
*/

func puissant(mot string) (estPuissant bool) {
	if len(mot) == 0 {
		return true
	}
	var pattern []string = []string{string(mot[0])}
	var found bool = false
	for i := 1; i < len(mot); i++ {
		fmt.Println(pattern)
		if i > len(mot)/2 && !found {
			return false
		}
		if found == false && pattern[0] == string(mot[i]) {
			var state bool = true
			for x := 1; x < len(pattern); x++ {
				if string(mot[i+x]) != pattern[x] {
					state = false
				}
			}
			if state {
				for x := 0; x < len(mot); x++ {
					if string(mot[x]) != pattern[x%len(pattern)] {
						state = false
					}
				}
				if state {
					found = true
				}
			}
			if !state {
				pattern = append(pattern, string(mot[i]))
			}
		} else {
			pattern = append(pattern, string(mot[i]))
		}
	}
	return true
}
