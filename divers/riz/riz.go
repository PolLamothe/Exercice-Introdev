package riz

/*
Dans un tableau, on place un grain de riz dans la première case, deux grains de riz dans la deuxième, quatre dans la troisième et ainsi de suite (2 fois plus de grains de riz dans chaque case que dans la précédente). On souhaite connaître le nombre total de grains de riz qu'il y aura dans un tableau en fonction de sa longueur (c'est-à-dire en fonction de son nombre de cases).

On considérera que les longueurs de tableaux prises dans le tests seront toujours assez petites pour que le nombre de grains de riz puisse être représenté par un entier en Go.

# Entrée
- longueur : le nombre de cases du tableau

# Sortie
- num : le nombre de grains de riz qui seront placés dans ce tableau

# Info
2023-2024, test 3, exercice 5
*/

func recur(longueur int) (sum, before int) {
	if longueur == 1 {
		return 1, 1
	}
	sum, before = recur(longueur - 1)
	return sum + 2*before, 2 * before
}

func riz(longueur int) (num int) {
	sum, _ := recur(longueur)
	return sum
}
