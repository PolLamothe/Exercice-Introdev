package prefixes

/*
La fonction numPrefixes doit compter le nombre de chaînes de caractères dans un
tableau qui commencent par une chaîne donnée.

# Entrées
- t : un tableau de chaînes de caractères
- s : une chaîne de caractères

# Sorties
- n : le nombre de chaînes de t qui commencent par s

# Exemple
numPrefixes([]string{"bonjour", "bonsoir", "salut", "bye bye"}, "bon") = 2
*/
func numPrefixes(t []string, s string) (n int) {
	if len(s) == 0 {
		return len(t)
	}
	var count int = 0
	for i := 0; i < len(t); i++ {
		if len(t[i]) > 0 {
			var state bool = true
			for x := 0; x < len(s); x++ {
				if t[i][x] != s[x] {
					state = false
				}
				if x == len(s)-1 && state {
					count++
				}
			}
		}
	}
	return count
}
