package palindrome

/*
Un palindrome est une chaîne de caractères qui se lit exactement pareil dans les deux sens (sans prendre en compte les espaces et les accents).
On peut aussi donner une définition récursive d'un palindrome (une fois les espaces et accents retirés d'une chaîne) :
la chaîne xsy (avec x, y des lettres et s une chaîne) est un palindrome si et seulement si x = y et s est un palindrome.
La fonction estPalindrome doit implanter cette définition récursive.

# Entrée
- s : une chaîne de caractère (sans espaces, sans accents, différente de nil)

# Sortie
- b : un booléen indiquant si s est un palindrome ou pas

# Exemples
estPalindrome("bonjour") = false
estPalindrome("kayak") = true
*/

func recur(s string, indexLeft, indexRight int) bool {
	if s[indexLeft] == ' ' {
		indexLeft++
	}
	if s[indexRight] == ' ' {
		indexRight--
	}
	if indexLeft == indexRight {
		return true
	}
	if s[indexLeft] == s[indexRight] {
		return recur(s, indexLeft+1, indexRight-1)
	} else {
		return false
	}
}

func estPalindrome(s string) (b bool) {
	if len(s) == 0 {
		return true
	}
	return recur(s, 0, len(s)-1)
}
