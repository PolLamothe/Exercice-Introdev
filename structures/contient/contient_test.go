package contient

import "testing"

func Test0(t *testing.T) {
	e3 := element{valeur: 3}
	e2 := element{valeur: 2, suivant: &e3}
	e1 := element{valeur: 1, suivant: &e2}
	e0 := element{valeur: 0, suivant: &e1}
	l := liste{debut: &e0}

	if !l.contient(2) {
		t.Fail()
	}
}

func Test1(t *testing.T) {
	e3 := element{valeur: 3}
	e2 := element{valeur: 2, suivant: &e3}
	e1 := element{valeur: 1, suivant: &e2}
	e0 := element{valeur: 0, suivant: &e1}
	l := liste{debut: &e0}

	if l.contient(4) {
		t.Fail()
	}
}

// Ajoutés après le test

type test struct {
	contenu  []int
	v        int
	contient bool
}

var testSet []test = []test{
	{contenu: []int{0, 1, 2, 3}, v: 2, contient: true},
	{contenu: []int{0, 1, 2, 3}, v: 4, contient: false},
	{contenu: []int{1, 5, 2, 0, 9, 2, 7, 5, 6, 3}, v: 9, contient: true},
	{contenu: []int{1, 5, 2, 0, 9, 2, 7, 5, 6, 3}, v: 12, contient: false},
	{contenu: []int{}, v: 1, contient: false},
}

func TestGeneral(t *testing.T) {
	for numTest, aTester := range testSet {

		l1 := creerListe(aTester.contenu)
		l2 := creerListe(aTester.contenu)

		contient := l1.contient(aTester.v)
		if aTester.contient != contient || !egal(l1, l2) {
			t.Error("Erreur sur l'entrée", numTest, "de testSet")
		}
	}
}

func creerListe(contenu []int) (l liste) {
	var e *element
	for i := len(contenu) - 1; i >= 0; i-- {
		newE := element{valeur: contenu[i], suivant: e}
		e = &newE
	}
	l.debut = e
	return
}

func egal(l1, l2 liste) bool {
	return egalElement(l1.debut, l2.debut)
}

func egalElement(e1, e2 *element) bool {
	if e1 == nil && e2 == nil {
		return true
	}

	if e1 == nil || e2 == nil {
		return false
	}

	if e1.valeur != e2.valeur {
		return false
	}

	return egalElement(e1.suivant, e2.suivant)
}
