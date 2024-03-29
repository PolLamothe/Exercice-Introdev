package existe

import "os"

/*
La fonction existe doit dire si un fichier dont le nom est donné en paramètre
existe ou pas.

# Entrée
- fName : un nom de fichier

# Sortie
- ok : un booléen qui vaut true si le fichier de nom fName existe et false sinon
*/

func existe(fName string) (ok bool) {
	_, err := os.Stat(fName)
	if err == nil {
		return true
	}
	return false
}
