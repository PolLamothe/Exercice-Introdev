package riz

import "testing"

func Test(t *testing.T) {
	if riz(1) != 1 {
		t.Error("riz(1) vaut 1")
	}

	if riz(2) != 3 {
		t.Error("riz(2) vaut 3")
	}

	if riz(3) != 7 {
		t.Error("riz(3) vaut 7")
	}
}

// Ajouté après le test
func Test54431454(t *testing.T) {
	if riz(4) != 15 {
		t.Error("riz(4) vaut 15")
	}

	if riz(5) != 31 {
		t.Error("riz(5) vaut 31")
	}

	if riz(6) != 63 {
		t.Error("riz(6) vaut 63")
	}

	if riz(7) != 127 {
		t.Error("riz(7) vaut 127")
	}

	if riz(8) != 255 {
		t.Error("riz(8) vaut 255")
	}

	if riz(9) != 511 {
		t.Error("riz(9) vaut 511")
	}

	if riz(10) != 1023 {
		t.Error("riz(10) vaut 1023")
	}

	if riz(11) != 2047 {
		t.Error("riz(11) vaut 2047")
	}

	if riz(12) != 4095 {
		t.Error("riz(12) vaut 4095")
	}

	if riz(13) != 8191 {
		t.Error("riz(13) vaut 8191")
	}

	if riz(14) != 16383 {
		t.Error("riz(14) vaut 16383")
	}
}
