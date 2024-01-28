package motdepasse

import "testing"

func Test(t *testing.T) {
	if estSolide("12345678") {
		t.Error("12345678 ne devrait pas être considéré comme un mot de passe solide")
	}

	if !estSolide("1Aabbbbb") {
		t.Error("1Aabbbbb devrait être considéré comme un mot de passe solide")
	}
}

// Ajouté après le test machine
func Test2(t *testing.T) {
	solides := []string{
		"1234ABcd",
		"1234ABcdetnausitenauistenauitsensuaite",
		"AAAAAAa1",
		"1Aabbbbb",
		"11111111AAAAAAAAaaaaaaaa",
		"zZ90aaaa",
	}

	for _, s := range solides {
		if !estSolide(s) {
			t.FailNow()
		}
	}

	pasSolides := []string{
		"1234ABa",
		"1234ABCD",
		"1234abcd",
		"ABCDabcd",
		"AAAAAAAAAAAAAAAAAAAAAAAAAA",
		"11111111111111111111111111",
		"aaaaaaaaaaaaaaaaaaaaaaaaaa",
		"aA1",
		"aaaaaaaaAAAAAAAA",
		"aaaaaaaa11111111",
		"AAAAAAAA11111111",
	}

	for _, s := range pasSolides {
		if estSolide(s) {
			t.FailNow()
		}
	}

}
