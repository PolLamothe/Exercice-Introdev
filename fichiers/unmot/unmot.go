package unmot

import (
	"errors"
	"io/ioutil"
)

/*
Étant donné un nom de fichier, on souhaite obtenir le premier mot du fichier ou, si ce n'est pas possible (le fichier n'existe pas ou ne contient pas de mot), une erreur. Pour simplifier, on considérera qu'il n'existe que deux sortes de fichiers :
- des fichiers contenant exactement un mot sur leur première ligne (et éventuellement d'autres lignes ensuite)
- des fichiers ne contenant aucun mot

# Entrée
- fName : le nom du fichier à traiter

# Sorties
- mot : le premier mot du fichier fName s'il existe (n'import quoi sinon)
- err : nil s'il existe un mot dans le fichier fName, errImpossible sinon

# Info
2022-2023, test3, exercice 5
*/

var errImpossible error = errors.New("Le fichier n'existe pas ou il ne contient aucun mot")

func premiermot(fName string) (mot string, err error) {
	data, err := ioutil.ReadFile(fName)
	if err != nil {
		return "", errImpossible
	}
	result := ""
	for i := 0; i < len(data); i++ {
		if len(result) > 0 && data[i] == '\n' {
			return result, nil
		} else if data[i] != '\n' {
			result += string(data[i])
		}
	}
	if len(result) > 0 {
		return result, nil
	} else {
		return result, errImpossible
	}
}
