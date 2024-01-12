package acrostiche

import (
	"fmt"
	"io/ioutil"
)

/*
La fonction acrostiche doit reconstituer le mot formé par le premier caractère
de chaque ligne d'un fichier, en ignorant les lignes vides.

# Entrée
- fName : le nom d'un fichier

# Sortie
- mot : la chaîne de caractère obtenue en mettant l'une après l'autre, dans
        l'ordre des lignes du fichier, les premières lettres de chacunes des
        lignes du fichier dont le nom est fName (on considère que ce fichier
        existe toujours)
*/

func acrostiche(fName string) (mot string) {
	data, error := ioutil.ReadFile(fName)
	if error != nil {
		return ""
	}
	var result string = ""
	var switche bool = true
	for i := 0; i < len(data); i++ {
		if switche && string(data[i]) != "\n" {
			result += string(data[i])
			switche = false
		}
		if string(data[i]) == "\n" {
			switche = true
		}
	}
	fmt.Println(result)
	return result
}
