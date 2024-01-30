package sortlist

/*
On dispose d'une structure de liste chaînée. Étant donnée une liste l on souhaite obtenir une nouvelle liste sortedL contenant les mêmes valeurs mais triées en ordre croissant. Ceci doit être fait sans modifier l. C'est le travail de la méthode Sort.

# Entrée
- l : une liste de valeurs

# Sortie
- sortedL : une liste contenant les mêmes valeurs que l mais triées en ordre croissant

# Info
2023-2024, test 3, exercice 8
*/

type List struct {
	Content *Element
}

type Element struct {
	V    int
	Next *Element
}

func tri(tab []int) []int {
	for i := 1; i < len(tab); i++ {
		if tab[i-1] > tab[i] {
			tab[i-1], tab[i] = tab[i], tab[i-1]
			return tri(tab)
		}
	}
	return tab
}

func getAll(l List) (tab []int) {
	e := l.Content
	for e != nil {
		tab = append(tab, e.V)
		e = e.Next
	}
	return
}

func (l List) Sort() (sortedL List) {
	tab := getAll(l)
	tab = tri(tab)
	sortedL = List{}
	for i := len(tab) - 1; i >= 0; i-- {
		sortedL.Content = &Element{V: tab[i], Next: sortedL.Content}
	}
	return
}
