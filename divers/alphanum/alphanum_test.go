package alphanum

import "testing"

func Test0(t *testing.T) {
	if !alphanum("Bonjour") {
		t.Fail()
	}
}

func Test1(t *testing.T) {
	if !alphanum("125") {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	if !alphanum("Sephirothdu44") {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	if alphanum("Il n'y a pas d'apostrophes ou d'espaces dans une chaîne alphanumérique (et pas d'accents non plus)") {
		t.Fail()
	}
}

// Ajoutés après le test

type test struct {
	s          string
	isAlphanum bool
}

var testSet []test = []test{
	{s: "Bonjour", isAlphanum: true},
	{s: "125", isAlphanum: true},
	{s: "Sephirothdu44", isAlphanum: true},
	{s: "Il n'y a pas d'apostrophes ou d'espaces dans une chaîne alphanumérique (et pas d'accents non plus)", isAlphanum: false},
	{s: "Bonjour ", isAlphanum: false},
	{s: " Bonjour", isAlphanum: false},
	{s: "", isAlphanum: true},
}

func TestGeneral(t *testing.T) {
	for numTest, aTester := range testSet {
		isAlphanum := alphanum(aTester.s)
		if aTester.isAlphanum != isAlphanum {
			t.Error("Erreur sur l'entrée", numTest, "de testSet")
		}
	}
}
