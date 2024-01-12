package nombrespairs

/*
La fonction sommeNombresPairs doit retourner la somme de tous les nombres
pairs contenus entre deux entiers (inclus).

# Entrées
- a : un nombre entier, une des bornes
- b : un nombre entier, l'autre borne

# Sortie
- somme : la somme de tous les nombres entiers compris entre a et b inclus

# Exemple
sommeNombresPairs(2, 7) = 12
*/
func sommeNombresPairs(a, b int) (somme int) {
	var somme1 int = 0
	if b < a && a >= 0 {
		var temp int = a
		a = b
		b = temp
	}
	for i := a; i <= b; i++ {
		if i%2 == 0 {
			somme1 += i
		}
	}
	return somme1
}
