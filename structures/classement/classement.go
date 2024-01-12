package classement

import (
	"strconv"
	"strings"
)

/*
La fonction classement doit, étant donné un tableau d'étudiants (supposé trié
par ordre de moyenne, de nom de famille en cas d'égalité, de prénom en cas de
nouvelle égalité — voir l'exercice "classer") former une chaîne de caractères
qui permettra un affichage correct du classement.

Chaque étudiant doit avoir son nom et son prénom affichés sur une ligne,
précédés de son classement, et ceci dans l'ordre du tableau. Si deux étudiants
ont la même moyenne, ils doivent avoir le même classement devant leurs noms.
Le classement d'un étudiant est le nombre d'étudiants classés avant lui + 1.

# Entrée
- classe : un tableau d'étudiants, trié

# Sortie
- affichage : une chaîne de caractères permettant l'affichage du classement

# Exemple
classement([]etudiant{
		{"Berdjugin", "Jean-François", 16.6},
		{"Hadj-Rabia", "Nassim", 15.2},
		{"Tamzalit", "Dalila", 15.2},
		{"Jezequel", "Loïg", 12.3},
		{"Jezequel", "Maël", 12.3},
	})

devrait permettre d'afficher le classement suivant :

1. Berdjugin Jean-François
2. Hadj-Rabia Nassim
2. Tamzalit Dalila
4. Jezequel Loïg
4. Jezequel Maël

# Info
2021-2022, test3, exercice 6
*/

type etudiant struct {
	nom, prenom string
	moyenne     float64
}

func classement(classe []etudiant) (affichage string) {
	var result []string
	var rank int = 0
	for i := 0; i < len(classe); i++ {
		if i > 0 {
			if classe[i].moyenne == classe[i-1].moyenne {
				rank++
			} else {
				rank = 0
			}
		}
		result = append(result, strconv.Itoa(i-rank+1))
		result = append(result, ". ")
		result = append(result, classe[i].nom)
		result = append(result, " ")
		result = append(result, classe[i].prenom)
		result = append(result, "\n")
	}
	return strings.Join(result, "")
}
