package classer

import "fmt"

/*
La fonction classer doit trier un tableau d'étudiants (tels que définis) ci-dessous
de celui qui a la meilleur moyenne (la plus élevée) à celui qui a la plus mauvaise
moyenne. En cas d'égalité de moyenne entre deux étudiants, celui qui a le nom de
famille qui arrive en premier dans l'ordre alphabétique doit être classer avant
l'autre (on peut utiliser <, >, <=, >=, == pour comparer les chaînes de caractères
par ordre alphabétique). Si deux étudiants ont la même moyenne et le même nom, on
met en premier celui qui a le prénom qui est en premier dans l'ordre alphabétique.

# Entrée
- promo : le tableau d'étudiants à trier

# Sortie
- classement : un tableau contenant les mêmes étudiants mais trié

# Info
2021-2022, test3, exercice 5
*/

type etudiant struct {
	nom, prenom string
	moyenne     float64
}

func decalRight(tab []etudiant, index int) []etudiant {
	var result []etudiant
	for i := 0; i < len(tab); i++ {
		if i == index {
			result = append(result, etudiant{})
		}
		result = append(result, tab[i])
	}
	return result
}

func supp(tab []etudiant, index int) []etudiant {
	var reusult []etudiant
	for i := 0; i < len(tab); i++ {
		if i != index {
			reusult = append(reusult, tab[i])
		}
	}
	return reusult
}

func classer(promo []etudiant) (classement []etudiant) {
	fmt.Println(promo)
	for i := len(promo) - 1; i > 0; i-- {
		if promo[i].moyenne > promo[i-1].moyenne {
			var temp etudiant = promo[i]
			promo = supp(promo, i)
			promo = decalRight(promo, i-1)
			promo[i-1] = temp
			return classer(promo)
		} else if promo[i].moyenne == promo[i-1].moyenne {
			if promo[i].nom < promo[i-1].nom {
				var temp etudiant = promo[i]
				promo = supp(promo, i)
				promo = decalRight(promo, i-1)
				promo[i-1] = temp
				return classer(promo)
			} else if promo[i].nom == promo[i-1].nom {
				if promo[i].prenom < promo[i-1].prenom {
					var temp etudiant = promo[i]
					promo = supp(promo, i)
					promo = decalRight(promo, i-1)
					promo[i-1] = temp
					return classer(promo)
				}
			}
		}
	}
	return promo
}
