package motdepasse

/*
On considère qu'un mot de passe est assez solide s'il contient au moins une lettre minuscule, une lettre majuscule, un chiffre et si, de plus, il fait au moins 8 caractères. La fonction estSolide doit indiquer si un mot de passe est solide ou pas.

#Entrée
- mdp : le mot de passe à tester

# Sortie
- solide : true si mdp est un mot de passe assez solide, false sinon

# Info
2023-2024, test 3, exercice 1
*/

var minuscules string = "abcdefghijklmnopqrstuvwxyz"
var majuscules string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var chiffres string = "0123456789"

func estSolide(mdp string) (solide bool) {
	state := false
	for i := 0; i < len(minuscules); i++ {
		for x := 0; x < len(mdp); x++ {
			if mdp[x] == minuscules[i] {
				state = true
			}
		}
	}
	if !state {
		return false
	}
	state = false
	for i := 0; i < len(majuscules); i++ {
		for x := 0; x < len(mdp); x++ {
			if mdp[x] == majuscules[i] {
				state = true
			}
		}
	}
	if !state {
		return false
	}
	state = false
	for i := 0; i < len(chiffres); i++ {
		for x := 0; x < len(mdp); x++ {
			if mdp[x] == chiffres[i] {
				state = true
			}
		}
	}
	if !state {
		return false
	}
	if len(mdp) < 8 {
		return false
	}
	return true
}
