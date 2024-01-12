package alphabet

/*
La fonction alphabet doit indiquer si deux mots sont dans l'ordre alphabétique.
On considérera que les mots en questions ne sont constitués que de lettres
minuscules non accentuées, donc prises dans l'ensemble {a, b, c, d, e, f, g, h,
i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z}.

# Entrées
- m1 : un mot constitué de lettres minuscules non accentuées
- m2 : un mot constitué de lettres minuscules non accentuées

# Sorties
- m1avantm2 : un booléen qui vaut true si m1 est (strictement) avant m2 dans
l'ordre alphabétique et qui vaut false sinon

# Indication
On peut comparer les caractères d'une chaîne de caractères comme on compare des
entiers. Ainsi, m1[i] < m2[j] vaut true si et seulement si le i-ième caractère
de m1 est avant le j-ième caractère de m2 dans l'alphabet.

# Info
2021-2022, test 1, exercice 7
*/

func boucle(m1 string, m2 string, alphabet []string) bool {
	var m1Count int
	var m2Count int
	for i := 0; i < len(m1); i = i + 1 {
		for x := 0; x < len(alphabet); x = x + 1 {
			if alphabet[x] == string(m1[i]) {
				m1Count = x
			}
			if alphabet[x] == string(m2[i]) {
				m2Count = x
			}
		}
		if m1Count < m2Count {
			return true
		} else if m2Count < m1Count {
			return false
		}
		if i == len(m1)-1 {
			return true
		}
		if i == len(m2)-1 {
			return false
		}
	}
	return true
}

func alphabet(m1, m2 string) (m1avantm2 bool) {
	if m1 == m2 {
		return false
	}
	if m1 == "" {
		return true
	}
	if m2 == "" {
		return false
	}
	var alphabet []string = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	return boucle(m1, m2, alphabet)
}
