package moitie

/*
La fonction moitieDePositifs indique si au moins de la moitié des nombres contenus dans un tableau sont strictement supérieurs à 0

# Entrée
- t : un tableau d'entiers

# Sortie
- reponse : un booléen qui vaut true si au moins la moitié des nombres de t sont strictement supérieurs à 0 et false sinon

# Info
2022-2023, test 1, exercice 6
*/

func moitieDePositifs(t []int) (reponse bool) {
	var count int = 0
	var negcount int = 0
	for i := 0; i < len(t); i++ {
		if t[i] > 0 {
			count++
		} else {
			negcount++
		}
	}
	return count >= negcount
}
