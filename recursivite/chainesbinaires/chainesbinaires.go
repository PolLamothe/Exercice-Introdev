package chainesbinaires

/*
Une chaîne binaire est une chaîne de caractères composée uniquement de 0 et de 1. On s'intéresse ici aux chaînes binaires sans 1 consécutifs. On peut générer de tels chaînes récursivement en considérant que
- si on dispose d'une chaîne s qui finit par 0 on peut l'étendre par 1 ou par 0, et
- si on dispose d'une chaîne s qui finit par 1 on ne peut l'étendre que par 0.
La fonction calculeChaines utilisera ce procédé (appliqué à chaque chaîne d'un tableau de chaînes) pour engendrer toutes les chaînes binaires sans 1 consécutifs d'une longueur donnée.

# Entrée
- n : un entier indiquant la longueur des chaînes à engendrer.

# Sortie
- chaines : un tableau de string contenant toutes les chaînes binaires de longueur n sans 1 consécutifs.

# Exemple
calculeChaines(3) = [000 001 010 100 101] (l'ordre n'a pas d'importance)
*/

func replaceAtIndex(str string, replacement rune, index int) string {
	var result string
	for i := 0; i < len(str); i++ {
		if i == index {
			result += string(replacement)
		} else {
			result += string(str[i])
		}
	}
	return result
}

func isInArray(index int, array []int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == index {
			return true
		}
	}
	return false
}

func pop(array []string, index []int) []string {
	var result []string
	for i := 0; i < len(array); i++ {
		if !isInArray(i, index) {
			result = append(result, array[i])
		}
	}
	return result
}

func recur(current string) []string {
	var zeroIndex int = -1
	for i := 0; i < len(current); i++ {
		if current[i] == '0' {
			zeroIndex = i
			current = replaceAtIndex(current, '1', i)
			break
		}
	}
	if zeroIndex == -1 {
		return []string{}
	}
	for i := 0; i < zeroIndex; i++ {
		current = replaceAtIndex(current, '0', i)
	}
	return append(recur(current), current)
}

func calculeChaines(n int) (chaines []string) {
	if n < 0 {
		return []string{}
	}
	var init string = ""
	for i := 0; i < n; i++ {
		init += "0"
	}
	var result []string = append(recur(init), init)
	var banIndex []int = []int{}
	for i := 0; i < len(result); i++ {
		for x := 1; x < len(result[i]); x++ {
			if result[i][x] == '1' && result[i][x-1] == '1' {
				banIndex = append(banIndex, i)
			}
		}
	}
	result = pop(result, banIndex)
	return result
}
