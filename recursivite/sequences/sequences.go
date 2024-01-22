package sequences

/*
Étant donnés deux valeurs k et n on souhaite calculer toutes les séquences strictement croissantes de k entiers en utilisant uniquement des entiers entre 1 et n. C'est ce que fait la fonction sequences.

Pour cet exercice, utiliser des boucles for réduira votre note (l'exercice rapportera 1 point s'il est fait avec boucles et 2 points s'il est fait sans boucles).

# Entrées
- k : un entier positif ou nul
- n : un entier supérieur ou égal à k

# Sortie
- seq : un tableau de tableau d'entiers. Chacun des tableaux obtenus doit être de longueur K, trié en ordre strictement croissant et contenir uniquement des nombres entre 1 et n. Seq doit contenir tous les tableaux possibles qui vérifient les conditions précédentes, et sans contenir de doublons. L'ordre des tableaux dans seq n'a pas d'importance.

# Exemples
sequences(2, 3) = [[1 2] [1 3] [2 3]]
sequences(3, 3) = [[1 2 3]]

# Info
2023-2024, test 2, exercice 9
*/

func recur(E, previous []int, length int) [][]int {
	if len(E) == 0 {
		return [][]int{}
	}
	var result [][]int = [][]int{}
	if len(previous) == length-1 {
		for i := 0; i < len(E); i++ {
			temp2 := append([]int(nil), previous...)
			temp := append(temp2, E[i])
			result = append(result, temp)
		}
	} else {
		for i := 0; i < len(E); i++ {
			result2 := recur(E[i+1:], append(previous, E[i]), length)
			result = append(result, result2...)
		}
	}

	return result
}

func sequences(k, n int) (seq [][]int) {
	var E []int = []int{}
	for i := 1; i <= n; i++ {
		E = append(E, i)
	}
	seq = recur(E, []int{}, k)
	return seq
}
