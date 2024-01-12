package alphanum

/*
On appelle caratères alphanumériques les lettres minuscules, les lettres majuscules et les chiffres. On souhaite vérifier qu'une chaîne de caractères est constituée uniquement de caractères alphanumériques. C'est le travail de la fonction alphanum.

# Entrée
- s : une chaîne de caractères

# Sortie
- isAlphanum : un booléen qui vaut true si s est entièrement constituée de caractères alphanumériques et qui vaut false sinon

# Info
2023-2024, test 2, exercice 3
*/

var alphaList []rune = []rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
	'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
	'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
	'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
	'U', 'V', 'W', 'X', 'Y', 'Z',
}

var numList []rune = []rune{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
}

func alphanum(s string) (isAlphanum bool) {
	for i := 0; i < len(s); i++ {
		var state bool = false
		for x := 0; x < len(alphaList); x++ {
			if string(alphaList[x]) == string(s[i]) {
				state = true
			}
		}
		for x := 0; x < len(numList); x++ {
			if string(numList[x]) == string(s[i]) {
				state = true
			}
		}
		if !state {
			return state
		}
	}
	return true
}
