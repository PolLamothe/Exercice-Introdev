package moitie

/*
La fonction moitieDePairs indique si au moins de la moitié des nombres contenus dans un tableau sont pairs

# Entrée
- t : un tableau d'entiers

# Sortie
- reponse : un booléen qui vaut true si au moins la moitié des nombres de t sont pairs et false sinon

# Info
2022-2023, test 4, exercice 4
*/

func moitieDePairs(t []int) (reponse bool) {
	var count int = 0
	for i := 0; i < len(t); i++ {
		if t[i]%2 == 0 {
			count++
		}
	}
	return float64(count) >= float64(len(t))/float64(2)
}
