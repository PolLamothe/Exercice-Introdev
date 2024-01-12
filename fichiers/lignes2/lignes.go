package lignes

import (
	"fmt"
	"io/ioutil"
)

/*
La fonction lignes doit compter le nombre de lignes non vides dans un
fichier dont le nom est indiqué en paramètre.

# Entrée
- fName : une chaîne de caractères correspondant à un nom de fichier

# Sortie
- nLignes : un entier indiquant le nombre de lignes non vides dans le
fichier de nom fName ou -1 si le fichier n'existe pas
*/

func lignes(fName string) (nLignes int) {
	data, error := ioutil.ReadFile(fName)
	if error != nil {
		return -1
	}
	count := 0
	for i := 1; i < len(data); i++ {
		if string(data[i]) == "\n" && string(data[i-1]) != "\n" {
			count++
		}
	}
	fmt.Println(count)
	return count
}
