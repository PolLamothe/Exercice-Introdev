package stevej

import (
	"errors"
	"fmt"
)

var errPasUneDate error = errors.New("la date indiquée n'est pas valide")

/*
Steve Jobs est né le 24 février 1955. La fonction ecart devra indiquer le nombre
de jours de différence entre une date (à partir du 1er janvier 1900 jusqu'au 31
décembre 2021) et la date de naissance de Steve Jobs. L'écart en jours sera toujours
un nombre positif ou nul (on ne se préoccupera pas de savoir si la date est antérieure
ou postérieure à la naissance de steve Jobs).

# Entrées
- j : un numéro de jour
- m : un numéro de mois
- a : un numéro d'année

# Sorties
- ej : l'écart en jours entre j/m/a et le 24/2/1955
- err : une erreur qui vaudra errPasUneDate quand j/m/a n'est pas une date valide et nil sinon

# Exemple
ecart(17, 4, 1986) = 11375
*/
func ecart(j, m, a uint) (ej uint, err error) {
	fmt.Println(j, m, a)
	if j == 0 || m == 0 {
		return 0, errPasUneDate
	}
	if a < 1900 || a > 2021 {
		return 0, errPasUneDate
	}
	if j > 31 {
		return 0, errPasUneDate
	}
	if m > 12 {
		return 0, errPasUneDate
	}
	if (m%2 == 0 && m < 7) && j > 30 {
		return 0, errPasUneDate
	}
	if (m%2 == 1 && m >= 7) && j > 30 {
		return 0, errPasUneDate
	}
	if (m%2 == 1 && m < 7) && j > 31 {
		return 0, errPasUneDate
	}
	if m == 2 && j > 29 {
		return 0, errPasUneDate
	}
	if m == 2 && j == 29 && a%4 != 0 {
		return 0, errPasUneDate
	}
	if a > 1956 {
		future, _ := ecart(j, m, a-1)
		if a%4 == 0 {
			return future + 366, nil
		} else {
			return future + 365, nil
		}
	}
	if a == 1956 {
		if m > 2 {
			future, _ := ecart(j, m, a-1)
			return future + 366, nil
		} else if m == 2 && j >= 24 {
			future, _ := ecart(j, m, a-1)
			return future + 366, nil
		} else if m == 2 {
			future, _ := ecart(j, m-1, a)
			return future + 31, nil
		} else if j > 1 {
			future, _ := ecart(j-1, m, a)
			return future + 1, nil
		} else {
			future, _ := ecart(31, 12, a-1)
			return future + 1, nil
		}
	}
	if a == 1955 {
		if m > 2 {
			if ((m%2 == 1 && m <= 7) || (m%2 == 0 && m > 7)) && m != 3 {
				future, _ := ecart(30, m-1, a)
				return future + j, nil
			} else if m != 3 {
				future, _ := ecart(31, m-1, a)
				return future + j, nil
			} else {
				future, _ := ecart(28, m-1, a)
				return future + j, nil
			}
		} else if m == 2 && j > 24 {
			future, _ := ecart(j-1, m, a)
			return future + 1, nil
		} else if m == 2 && j == 24 {
			return 0, nil
		} else if m == 2 && j < 24 {
			future, _ := ecart(j, m+1, a)
			future -= 28 - 25 + 1
			return future, nil
		} else if m == 1 {
			if j >= 29 {
				future, _ := ecart(j-1, m, a)
				future--
				return future, nil
			}
			if j > 24 {
				future, _ := ecart(j, m+1, a)
				future += 31
				return future, nil
			}
			if j <= 24 {
				future, _ := ecart(j+1, m, a)
				future--
				return future, nil
			}
		}
	} else {
		future, _ := ecart(j, m, a+1955-a)
		future -= -(1955 - a) * 365
		return future, nil
	}
	return 0, nil
}
