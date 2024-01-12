package livres

/*
On dispose d'une structure de données permettant de réprésenter des livres avec un titre, un prix et un nombre de page.

On souhaite trier un tableau de livres soit par prix (du moins cher au plus cher) soit par nombre de pages (du plus grand nombre de pages au plus petit nombre de pages). Dans les deux cas, si deux livres ne sont pas distinguables, on les triera par ordre alphabétique de leur titre (en utilisant l'opérateur < sur les chaînes de caractères). La fonction trier doit réaliser ce tri en place.

# Entrée
- tab : le tableau de livres à trier
- parPrix : un booléen indiquant le tri à réaliser, s'il vaut true on triera les livres du moins cher au plus cher, s'il vaut false on triera les livres de celui qui a le plus de pages à celui qui en a le moins.

# Info
2022-2023, test3, exercice 4
*/

type livre struct {
	titre    string
	prix     float64
	numPages int
}

func trier(tab []livre, parPrix bool) []livre {
	if parPrix {
		for i := 0; i < len(tab)-1; i++ {
			if tab[i].prix > tab[i+1].prix {
				var temp livre = tab[i]
				tab[i] = tab[i+1]
				tab[i+1] = temp
				return trier(tab, true)
			} else if tab[i].prix == tab[i+1].prix {
				if tab[i].titre > tab[i+1].titre {
					var temp livre = tab[i]
					tab[i] = tab[i+1]
					tab[i+1] = temp
					return trier(tab, true)
				}
			}
		}
		return tab
	} else {
		for i := 0; i < len(tab)-1; i++ {
			if float64(tab[i].numPages) < float64(tab[i+1].numPages) {
				var temp livre = tab[i]
				tab[i] = tab[i+1]
				tab[i+1] = temp
				return trier(tab, false)
			} else if float64(tab[i].numPages) == float64(tab[i+1].numPages) {
				if tab[i].titre > tab[i+1].titre {
					var temp livre = tab[i]
					tab[i] = tab[i+1]
					tab[i+1] = temp
					return trier(tab, false)
				}
			}
		}
		return tab
	}
}
