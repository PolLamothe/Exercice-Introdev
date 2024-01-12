package unique

/*
Le but de cet exercice est de créer un algorithme permettant de trier une array contenant des entiers, le tout sans utiliser d'array supplémentaire, en modifiant celle reçu
*/

func tri(tabs []int) []int {
	for i := 1; i < len(tabs); i++ { //pour chaque élément du tableau en commencant par le deuxième
		if tabs[i] < tabs[i-1] { //si jamais cet élément est plus petit que celui d'avant
			var temp int = tabs[i-1]  //on sauvegarde le terme d'avant
			tabs[i-1] = tabs[i]       //on décale le terme actuelle a celui d'avant
			tabs[i] = temp            //on décale la valeur du terme d'avant a l'actuelle
			if tabs[i] != tabs[i-1] { //si jamais le terme d'avant était plus grand et pas égale (pour ne pas faire une boucle infini)
				return tri(tabs) //on recommence la vérification
			}
		}
	}
	return tabs
}

func unique(tabs []int) []int {
	return tri(tabs) //on appelle la fonction
}
